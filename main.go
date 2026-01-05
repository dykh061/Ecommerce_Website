package main

import (
	"fmt"
	"log"
	"os"

	"OpenMarket/component/appctx"
	"OpenMarket/component/uploadProvider"
	"OpenMarket/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect mysql: %v", err)
	}

	s3Provider := uploadprovider.NewS3Provider(
		os.Getenv("S3BucketName"),
		os.Getenv("S3Region"),
		os.Getenv("S3APIKey"),
		os.Getenv("S3SecretKey"),
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_PUBLIC_DOMAIN"),
	)

	appCtx := appctx.NewAppContext(
		db,
		s3Provider,
		os.Getenv("SYSTEM_SECRET"),
	)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(middleware.Recover(appCtx))

	// gắn routes
	setupRoutes(appCtx, r)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("server listening on %s", port)
	r.Run(port)
}
