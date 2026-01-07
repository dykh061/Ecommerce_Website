package main

import (
	"OpenMarket/component/appctx"
	"OpenMarket/middleware"
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
	// PRODUCT (PUBLIC)
	// =================================================
	v1.GET("/products", productgin.ListPublicProduct(appCtx))
	v1.GET("/products/:id", productgin.GetProductDetail(appCtx))
	v1.GET("/products/:id/variants", productgin.ListVariant(appCtx))
	v1.GET("/products/:id/galleries", productgin.GetImages(appCtx))
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
		sellerProduct.POST("", productgin.CreateProduct(appCtx))
		sellerProduct.PATCH("/:id", productgin.UpdateProduct(appCtx))
		sellerProduct.POST("/:id/variants", productgin.CreateVariant(appCtx))
		sellerProduct.PATCH("/:id/variant/:vid", productgin.UpdateVariant(appCtx))
		sellerProduct.POST("/:id/galleries", productgin.CreateProductGallery(appCtx))
	}

	// =================================================
	// UPLOAD
	// =================================================
	v1.POST(
		"/upload",
		middleware.RequiredAuthenHeader(appCtx),
		uploadgin.UpLoadImage(appCtx),
	)
}
