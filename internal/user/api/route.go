package api

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/internal/middleware"
	"ndy/realworld-gin/internal/user/app"
)

func Routes(g *gin.RouterGroup, l *app.Logic) {
	// middlewares
	um := middleware.JsonRoot("user", "user")
	umRespOnly := middleware.JsonRoot("", "user")
	am := middleware.Auth()

	// register routes
	g.POST("/users", um, RegisterHandler(l))
	g.GET("/user", am, umRespOnly, GetCurrentUserHandler(l))
	g.PUT("/user", am, um, UpdateUserHandler(l))
}
