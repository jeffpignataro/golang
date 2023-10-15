package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "go.uber.org/zap"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/alive", func(c *gin.Context) {
			log.S().Info("alive")
			c.JSON(200, gin.H{"message": "ok"})
		})
		v1.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ok"})
		})
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
		v1.GET("/quitquitquit", func(c *gin.Context) {
			log.S().Info("quitting application")
			c.JSON(200, gin.H{"message": "quitting application"})
			os.Exit(1)
		})
	}
	r.Run(":8080")
}
