package handlers

import (
	"net/http"
	"time"

	"github.com/ayanAhm4d/URL-shortener/models"
	"github.com/ayanAhm4d/URL-shortener/redis"
	"github.com/ayanAhm4d/URL-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShortenURL(c *gin.Context) {
	var req models.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil || !utils.IsValidURL(req.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request or URL"})
		return
	}

	req.URL = utils.EnforceHTTP(req.URL)
	shortID := req.CustomShort
	if shortID == "" {
		shortID = uuid.New().String()[:6]
	}

	if redis.Client.Get(redis.Ctx, shortID).Val() != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Custom short URL already exists"})
		return
	}

	expiry := 24 * time.Hour
	if req.Expiry > 0 {
		expiry = time.Duration(req.Expiry) * time.Hour
	}

	err := redis.Client.Set(redis.Ctx, shortID, req.URL, expiry).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url":  utils.GetBaseURL() + "/" + shortID,
		"expires_in": expiry.String(),
	})
}

func ResolveURL(c *gin.Context) {
	shortID := c.Param("short")
	url := redis.Client.Get(redis.Ctx, shortID).Val()

	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found or expired"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
