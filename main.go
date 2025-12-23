package main

import (
	"fmt"
	"log"
	"os"

	"OpenMarket/component/appctx"
	"OpenMarket/middleware"
	productgin "OpenMarket/module/product/transport/gin"
	ginseller "OpenMarket/module/seller/transport/gin"
	ginuser "OpenMarket/module/user/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect mysql: %v", err)
	}

	appContext := appctx.NewAppContext(db, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.POST("/users", ginuser.Register(appContext))
	r.POST("/authenticate", ginuser.Login(appContext))
	r.GET("/profile", middleware.RequiredAuthenHeader(appContext), ginuser.Profile(appContext))
	r.PUT("/users/:id", ginuser.UpdateUser(appContext))
	r.DELETE("/users/:id", ginuser.DeleteUser(appContext))
	r.POST("/products", productgin.CreateProduct(appContext))
	r.GET("/products", productgin.ListProduct(appContext))
	// r.DELETE("/products/:id", productgin.DeleteProduct(appContext))
	r.POST("/sellers", ginseller.CreateSeller(appContext))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("server listening on %s", port)
	r.Run(port)
}
