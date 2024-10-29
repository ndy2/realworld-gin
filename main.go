package main

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/auth"
	"ndy/realworld-gin/route"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		users := api.Group("/users")
		{
			users.POST("/login", route.HandleJsonRootMiddleware("user", "user"), auth.Authentication)
			//users.POST("/", route.HandleJsonRootMiddleware("user", "user"), auth.Registration)
		}
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
