package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func main() {
	r := gin.Default()
	r.GET("/ping", handler)

	r.Run()
}
