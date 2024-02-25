package main

import (
	"net/http"

	"github.com/AngelinaSantana/Docker-/configs"
	"github.com/AngelinaSantana/Docker-/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
