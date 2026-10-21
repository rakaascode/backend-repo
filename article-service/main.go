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
	r.Use(corsMiddleware())

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

// corsMiddleware allows browser clients to access this public API directly.
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
