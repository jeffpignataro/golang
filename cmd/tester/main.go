package main

import (
	"golang/pkg/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

var proxies = []string{"127.0.0.1"}

func main() {
	ws := gin.Default()
	ws.SetTrustedProxies(proxies)

	ws.GET("/alive", alive)
	ws.GET("/ping", ping)

	ws.Run(":8888")
}

func alive(c *gin.Context) {
	c.JSON(200, "Alive")
}

func ping(c *gin.Context) {
	r, _ := rest.Call("https://www.google.com", http.MethodGet, nil, nil, nil)
	c.JSON(200, string(r))
}
