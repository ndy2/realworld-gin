package profile

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/util"
	"net/http"
)

func GetProfileHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUsername, _ := c.Get("username")
		profileId, _ := c.Get("profileId")
		util.Log.Info("Getting profile", zap.String("username", currentUsername.(string)), zap.Int("profileId", profileId.(int)))

		resp, err := l.GetProfile(currentUsername.(string), profileId.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Got profile", zap.String("username", resp.Username))

		// Return the response
		c.Set("resp", resp)
	}
}

type GetProfileResponse profileResponse

type profileResponse struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}
