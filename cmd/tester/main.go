package main

import (
	"golang/pkg/rest"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	ws := gin.Default()
	ws.GET("/alive", alive)
	ws.GET("/ping", ping)
	ws.Run()
}

func alive(c *gin.Context) {
	c.JSON(200, "Alive")
}

func ping(c *gin.Context) {
	r, _ := rest.Call("https://www.google.com", http.MethodGet, nil, nil, nil)
	log.Info(r)
}
