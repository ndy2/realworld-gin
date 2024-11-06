package api

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/internal/middleware"
	"ndy/realworld-gin/internal/profile/app"
)

func Routes(g *gin.RouterGroup, l *app.Logic) {
	// middlewares
	pmRespOnly := middleware.JsonRoot("", "profile")
	oam := middleware.OptionalAuth()

	// register routes
	g.GET("/profiles/:username", oam, pmRespOnly, GetProfileHandler(l))
}
