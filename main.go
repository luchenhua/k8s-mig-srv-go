package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v0 := router.Group("/srv-go")
	{
		v0.GET("/ping", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, "pong")
		})
		v0.GET("/users", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, "GET")
		})
		v0.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.SecureJSON(http.StatusOK, "GET:"+id)
		})
		v0.POST("/users", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, "POST")
		})
		v0.PUT("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.SecureJSON(http.StatusCreated, "PUT: "+id)
		})
	}

	router.Run(":3000")
}
