package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Access-Control-Allow-Origin", "Content-Type", "X-CSRF-Token"}
	config.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(config))

	router.GET("/api/joke", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Knock knock..."})
	})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
