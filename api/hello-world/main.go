package main

import (
	"github.com/gin-gonic/gin"
	helloworld "github.com/jeffpignataro/golang/pkg/hello-world"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, helloworld.Hello())
	})
	r.Run()
}
