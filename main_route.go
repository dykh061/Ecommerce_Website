package main

import (
	"OpenMarket/component/appctx"
	"OpenMarket/middleware"
	admingin "OpenMarket/module/admin/transport/gin"
	cartgin "OpenMarket/module/cart/transport/gin"
	ordergin "OpenMarket/module/order/transport/gin"
	productgin "OpenMarket/module/product/transport/gin"

	ginseller "OpenMarket/module/seller/transport/gin"
	uploadgin "OpenMarket/module/upload/transport/gin"
	ginuser "OpenMarket/module/user/transport/gin"

	"github.com/gin-gonic/gin"
)

func setupRoutes(appCtx appctx.AppContext, r *gin.Engine) {

	// =======================
	// API v1
	// =======================
	v1 := r.Group("/v1")

	// =================================================
	// AUTH / USER
	// =================================================
	v1.POST("/auth/login", ginuser.Login(appCtx))
	v1.POST("/users", ginuser.Register(appCtx))

	userAuth := v1.Group(
		"/users",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{
		userAuth.GET("/profile", ginuser.Profile(appCtx))
		userAuth.PATCH("", ginuser.UpdateUser(appCtx))
		userAuth.PATCH("/password", ginuser.ChangePassword(appCtx))
		userAuth.DELETE("", ginuser.DeleteUser(appCtx))
	}

	// =================================================
	// ADDRESS (AUTH)
	// =================================================
	address := v1.Group(
		"/address",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{

		address.PATCH("/:id", ginuser.UpdateAddress(appCtx))
		address.DELETE("/:id", ginuser.DeleteAddress(appCtx))
		address.GET("", ginuser.GetListAddress(appCtx))
		address.POST("", ginuser.CreateAddress(appCtx))
	}

	// =================================================
	// SELLER (PUBLIC)
	// =================================================
	v1.GET("/sellers", ginseller.ListSeller(appCtx))
	v1.GET("/sellers/:id", ginseller.GetSeller(appCtx))

	// =================================================
	// SELLER (AUTH)
	// =================================================
	sellerAuth := v1.Group(
		"/sellers",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{
		sellerAuth.POST("", ginseller.CreateSeller(appCtx))
		sellerAuth.GET("/me", ginseller.GetMyShop(appCtx))
		sellerAuth.PATCH("", ginseller.UpdateSeller(appCtx))
		sellerAuth.DELETE("", ginseller.DeleteSeller(appCtx))
	}

	// =================================================
	// CATEGORY (PUBLIC)
	// =================================================
	v1.GET("/categories", productgin.ListCategories(appCtx))
	v1.GET("/categories/:category_id/attributes", productgin.GetCategoryAttributes(appCtx))

	// =================================================
	// PRODUCT (PUBLIC)
	// =================================================
	v1.GET("/products", productgin.ListPublicProduct(appCtx))
	v1.GET("/products/:id", productgin.GetProductDetail(appCtx))
	v1.GET("/products/:id/variants", productgin.ListVariant(appCtx))
	v1.GET("/products/:id/galleries", productgin.GetImages(appCtx))
	v1.GET("/products/:id/attributes", productgin.GetProductAttributes(appCtx))
	v1.POST("/products/variants", productgin.AdjustStock(appCtx))

	// =================================================
	// PRODUCT (SELLER)
	// =================================================
	sellerProduct := v1.Group(
		"/seller/products",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{
		sellerProduct.GET("", productgin.ListSellerProduct(appCtx))
		sellerProduct.GET("/:id", productgin.GetSellerProductDetail(appCtx))
		sellerProduct.POST("", productgin.CreateProduct(appCtx))
		sellerProduct.PATCH("/:id", productgin.UpdateProduct(appCtx))
		sellerProduct.DELETE("/:id", productgin.DeleteProduct(appCtx))

		// Variant routes
		sellerProduct.GET("/:id/variant/:vid", productgin.GetSellerVariantDetail(appCtx))
		sellerProduct.POST("/:id/variants", productgin.CreateVariant(appCtx))
		sellerProduct.POST("/:id/variants/check-duplicate", productgin.CheckVariantDuplicate(appCtx))
		sellerProduct.PATCH("/:id/variant/:vid", productgin.UpdateVariant(appCtx))
		sellerProduct.PATCH("/:id/variant/:vid/status", productgin.ToggleVariantStatus(appCtx))
		sellerProduct.DELETE("/:id/variant/:vid", productgin.DeleteVariant(appCtx))

		// Gallery routes
		sellerProduct.POST("/:id/galleries", productgin.CreateProductGallery(appCtx))
		sellerProduct.PATCH("/:id/galleries/:gallery_id/main", productgin.SetMainGallery(appCtx))
		sellerProduct.DELETE("/:id/galleries/:gallery_id", productgin.DeleteGallery(appCtx))
	}

	// =================================================
	// CART (AUTH)
	// =================================================
	cart := v1.Group(
		"/carts",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{
		cart.POST("/items", cartgin.AddToCart(appCtx))
		cart.PATCH("/items", cartgin.AdJustProduct(appCtx))
		cart.DELETE("/items", cartgin.RemoveItemFromCart(appCtx))
		cart.POST("", cartgin.CreateCart(appCtx)) // dùng test tạm thôi chứ không cần thiết
		cart.GET("", cartgin.GetListItem(appCtx))
		cart.GET("/items/detail", cartgin.GetListItemDetail(appCtx))

	}

	// ==================================================
	// ORDER (AUTH)
	// =================================================
	order := v1.Group(
		"/orders",
		middleware.RequiredAuthenHeader(appCtx),
	)
	{
		order.GET("/:order_id", ordergin.GetDetailOrder(appCtx))
		order.GET("", ordergin.GetListOrder(appCtx))
		order.POST("", ordergin.CreateOrder(appCtx))
		order.PATCH("", ordergin.CanCelOrder(appCtx))

	}

	// =================================================
	// UPLOAD
	// =================================================
	v1.POST(
		"/upload",
		middleware.RequiredAuthenHeader(appCtx),
		uploadgin.UpLoadImage(appCtx),
	)

	// =================================================
	// ADMIN (Staff/Admin Panel)
	// =================================================

	// Admin Auth (Public - no token required)
	v1.POST("/admin/auth/login", admingin.Login(appCtx))

	// DEV-ONLY seed endpoints (guarded by env + non-release mode)
	v1.POST("/dev/seed/create-staff", admingin.DevSeedCreateStaff(appCtx))

	// Admin routes (requires admin authentication)
	admin := v1.Group(
		"/admin",
		middleware.RequireAdminAuth(appCtx),
	)
	{
		// Staff profile
		admin.GET("/profile", admingin.Profile())

		// Seller management (requires moderator or admin role)
		adminSellers := admin.Group("/sellers", middleware.RequireModerator())
		{
			adminSellers.GET("", admingin.ListSellers(appCtx))
			adminSellers.GET("/:id", admingin.GetSeller(appCtx))
			adminSellers.PATCH("/:id/status", admingin.UpdateSellerStatus(appCtx))
			adminSellers.POST("/:id/lock", admingin.LockSeller(appCtx))
			adminSellers.POST("/:id/unlock", admingin.UnlockSeller(appCtx))
		}
	}
}
