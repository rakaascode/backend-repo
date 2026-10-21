package controllers

import (
	"net/http"
	"time"

	"article-service/config"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "UP",
		"service":   "article-service",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func HealthCheckDB(c *gin.Context) {
	// Ambil koneksi sql.DB dari GORM
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "DOWN",
			"database": "Failed to get DB instance",
			"error":    err.Error(),
		})
		return
	}

	// Ping database dengan timeout 2 detik
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "DOWN",
			"database": "Unreachable",
			"error":    err.Error(),
		})
		return
	}

	stats := sqlDB.Stats()

	c.JSON(http.StatusOK, gin.H{
		"status":    "UP",
		"timestamp": time.Now().Format(time.RFC3339),
		"database": gin.H{
			"status":           "UP",
			"open_connections": stats.OpenConnections,
			"in_use":           stats.InUse,
			"idle":             stats.Idle,
		},
	})
}
