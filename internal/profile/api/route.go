package api

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/internal/middleware"
	"ndy/realworld-gin/internal/profile/app"
)

func Routes(g *gin.RouterGroup, l *app.Logic) {
	// middlewares
	pmRespOnly := middleware.JsonRoot("", "profile")
	am := middleware.Auth()

	// register routes
	g.GET("/profiles/:username", am, pmRespOnly, GetProfileHandler(l))
}
