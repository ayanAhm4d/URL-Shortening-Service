package main

import (
	"github.com/ayanAhm4d/URL-shortener/config"
	"github.com/ayanAhm4d/URL-shortener/handlers"
	"github.com/ayanAhm4d/URL-shortener/middleware"
	"github.com/ayanAhm4d/URL-shortener/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	redis.InitRedis()

	r := gin.Default()
	r.Use(middleware.RateLimiter())

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:short", handlers.ResolveURL)

	r.Run(":8080")
}
