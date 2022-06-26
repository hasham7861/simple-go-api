package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Message: "pong",
		})
	})
	r.Run()
}
