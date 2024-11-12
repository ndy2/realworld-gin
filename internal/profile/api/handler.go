package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/profile/app"
	"ndy/realworld-gin/internal/util"
	"net/http"
)

func GetProfileHandler(l *app.Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetUsername := c.Param("username")
		util.Log.Info("Getting profile", zap.String("username", targetUsername))

		// Get the current user if authenticated
		var currentUserId, currentUserProfileId int
		var currentUsername string
		if c.GetBool("Authenticated") {
			currentUserId = c.GetInt("userId")
			currentUserProfileId = c.GetInt("profileId")
			currentUsername = c.GetString("username")
		}

		// Get the profile
		resp, err := (*l).GetProfile(currentUserId, currentUserProfileId, currentUsername, targetUsername)

		// Profile not found
		if errors.Is(err, app.ErrProfileNotFound) {
			util.Log.Info("Profile not found", zap.String("username", targetUsername))
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
			return
		}
		// Other errors
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Got profile", zap.String("username", resp.Username))

		// Return the response
		c.Set("resp", resp)
	}
}
