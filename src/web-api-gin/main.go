package main

import (
	"fmt"
	"golang/src/web-api-gin/entities"
	"golang/src/web-api-gin/handlers"
	"golang/src/web-api-gin/repositories"
	"golang/src/web-api-gin/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Web Api Gin")
	dsn := "root@tcp(127.0.0.1:3306)/web-api-gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error!")
	}

	err = db.AutoMigrate(&entities.Book{})

	if err != nil {
		log.Fatal("DB Automigration Error!")
	}

	bookRepository := repositories.NewRepository(db)
	bookService := services.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)

	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	err = router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
