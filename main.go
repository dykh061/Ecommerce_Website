package main

import (
	uploadprovider "OpenMarket/component/uploadProvider"
	uploadgin "OpenMarket/module/upload/transport/gin"

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

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3PublicDomain := os.Getenv("S3_PUBLIC_DOMAIN")
	s3Endpoint := os.Getenv("S3_ENDPOINT")

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

	s3Provider := uploadprovider.NewS3Provider(s3BucketName,
		s3Region,
		s3APIKey,
		s3SecretKey,
		s3Endpoint,
		s3PublicDomain,
	)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8MB
	r.Use(middleware.Recover(appContext))

	// User routes */
	r.POST("/users", ginuser.Register(appContext))
	r.POST("/authenticate", ginuser.Login(appContext))
	r.GET("/profile", middleware.RequiredAuthenHeader(appContext), ginuser.Profile(appContext))
	r.PUT("/users/:id", ginuser.UpdateUser(appContext))
	r.DELETE("/users/:id", ginuser.DeleteUser(appContext))

	// Product routes */
	r.POST("/products", productgin.CreateProduct(appContext))
	r.GET("/products", productgin.ListProduct(appContext))
	// r.DELETE("/products/:id", productgin.DeleteProduct(appContext))
	r.POST("/sellers", ginseller.CreateSeller(appContext))

	//upload
	r.POST("/upload", uploadgin.UpLoadImage(appContext))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("server listening on %s", port)
	r.Run(port)
}
