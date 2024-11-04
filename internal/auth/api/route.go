package api

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/middleware"
)

func Routes(g *gin.RouterGroup, l *app.Logic) {
	// middlewares
	um := middleware.JsonRoot("user", "user")

	// register routes
	g.POST("/users/login", um, AuthenticationHandler(l))
}
