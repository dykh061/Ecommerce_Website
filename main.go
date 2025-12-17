package main

import (
	"fmt"
	"log"
	"os"

	"OpenMarket/component/appctx"
	productgin "OpenMarket/module/product/transport/gin"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect mysql: %v", err)
	}

	// if err := db.AutoMigrate(
	// 	&usermodel.User{},
	// 	&productmodel.Product{},
	// 	// &sellermodel.Seller{},
	// ); err != nil {
	// 	log.Fatalf("cannot migrate: %v", err)
	// }

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.POST("/users", ginuser.CreateUser(appContext))
	r.POST("/products", productgin.CreateProduct(appContext))
	r.GET("/products", productgin.ListProduct(appContext))
	r.DELETE("/products/:id", productgin.DeleteProduct(appContext))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("server listening on %s", port)
	r.Run(port)
}
