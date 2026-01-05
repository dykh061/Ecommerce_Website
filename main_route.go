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
	// PUBLIC API - v1
	// =======================
	v1 := r.Group("/v1")

	// -------- USER --------
	v1.POST("/users", ginuser.Register(appCtx))
	v1.POST("/auth/login", ginuser.Login(appCtx))
	v1.GET(
		"/users/profile",
		middleware.RequiredAuthenHeader(appCtx),
		ginuser.Profile(appCtx),
	)
	v1.PUT(
		"/users",
		middleware.RequiredAuthenHeader(appCtx),
		ginuser.ChangePassword(appCtx),
	)
	v1.PUT(
		"/users/:id",
		middleware.RequiredAuthenHeader(appCtx),
		ginuser.UpdateUser(appCtx),
	)
	v1.DELETE(
		"/users/:id",
		middleware.RequiredAuthenHeader(appCtx),
		ginuser.DeleteUser(appCtx),
	)

	// -------- SELLER --------
	v1.POST(
		"/sellers",
		middleware.RequiredAuthenHeader(appCtx),
		ginseller.CreateSeller(appCtx),
	)
	v1.GET("/sellers", ginseller.ListSeller(appCtx))
	v1.GET(
		"/sellers/:id",
		ginseller.GetSeller(appCtx),
	)
	v1.PUT(
		"/sellers",
		middleware.RequiredAuthenHeader(appCtx),
		ginseller.UpdateSeller(appCtx),
	)
	v1.DELETE(
		"/sellers",
		middleware.RequiredAuthenHeader(appCtx),
		ginseller.DeleteSeller(appCtx),
	)
	v1.GET(
		"/sellers/my-shop",
		middleware.RequiredAuthenHeader(appCtx),
		ginseller.GetMyShop(appCtx),
	)

	// -------- PRODUCT --------

	v1.POST("/products",
		middleware.RequiredAuthenHeader(appCtx),
		productgin.CreateProduct(appCtx),
	)

	v1.POST("/products/:id/variant",
		middleware.RequiredAuthenHeader(appCtx),
		productgin.CreateVariant(appCtx),
	)

	// -------- GALLERY --------
	v1.POST("/products/:id/galleries",
		middleware.RequiredAuthenHeader(appCtx),
		productgin.CreateProductGallery(appCtx),
	)

	// -------- UPLOAD --------
	v1.POST(
		"/upload",
		middleware.RequiredAuthenHeader(appCtx),
		uploadgin.UpLoadImage(appCtx),
	)
}
