package main

import (
	"article-service/config"
	"article-service/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan...")
	}

	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/health", controllers.HealthCheck)
	r.GET("/health/db", controllers.HealthCheckDB)

	r.POST("/article/", controllers.CreatePost)
	r.GET("/article/:id/:offset", controllers.GetPosts)
	r.GET("/article/:id", controllers.GetPostByID)
	r.PUT("/article/:id", controllers.UpdatePost)
	r.DELETE("/article/:id", controllers.DeletePost)

	err := r.Run(":8081")
	if err != nil {
		return
	}

}
