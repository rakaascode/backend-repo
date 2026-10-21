package controllers

import (
	"net/http"
	"strconv"

	"article-service/config"
	"article-service/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully", "data": input})
}

func GetPosts(c *gin.Context) {
	// Gin requires the same wildcard name as /article/:id at this position.
	limitStr := c.Param("id")
	offsetStr := c.Param("offset")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	if limit == 0 {
		limit = 10
	}

	var posts []models.Post
	config.DB.Limit(limit).Offset(offset).Find(&posts)

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data
	config.DB.Model(&post).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully", "data": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
