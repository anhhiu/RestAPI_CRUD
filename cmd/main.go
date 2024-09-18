package main

import (
	"log"
	"net/http"

	"RestApi_CRUD/internal/connect"
	"RestApi_CRUD/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối tới cơ sở dữ liệu
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Khởi tạo Gin router
	r := gin.Default()

	// Định nghĩa các route
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", handler.CreateItem(db))
			items.GET("", handler.GetAllItems(db))
			items.GET("/:id", handler.GetItemByID(db))
			items.PUT("/:id", handler.UpdateItem(db))
			items.DELETE("/:id", handler.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
