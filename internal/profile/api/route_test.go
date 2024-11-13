package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/internal/profile/app"
	"ndy/realworld-gin/internal/profile/dto"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var r *gin.Engine
var g *gin.RouterGroup
var w *httptest.ResponseRecorder
var ctx *gin.Context

var mockLogic *app.MockLogic

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	g = r.Group("/api")
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)

	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()
	mockLogic = app.NewMockLogic(ctrl)

	code := m.Run()

	os.Exit(code)
}

func TestRoutes_Profiles_Username_Success_Authenticated(t *testing.T) {
	// Given
	resp := dto.GetProfileResponse{
		Username:  "testuser",
		Bio:       "test bio",
		Image:     "test image",
		Following: false,
	}
	mockLogic.EXPECT().GetProfile(1, 1, "testuser", "targetuser").Return(resp, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"GET",
		"/api/profiles/targetuser",
		nil,
	)
	ctx.Request = httpReq
	ctx.Set("authenticated", true)
	ctx.Set("userId", 1)
	ctx.Set("profileId", 1)
	ctx.Set("username", "testuser")

	r.ServeHTTP(w, httpReq)

	// Then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{
		"profile": {
			"username": "testuser",
			"bio": "test bio",
			"image": "test image",
			"following": false
		}
	}`, w.Body.String())
}

func TestRoutes_Profiles_Username_Success_Unauthenticated(t *testing.T) {
	// Given
	resp := dto.GetProfileResponse{
		Username:  "targetuser",
		Bio:       "target bio",
		Image:     "",
		Following: false,
	}
	mockLogic.EXPECT().GetProfile(0, 0, "", "targetuser").Return(resp, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"GET",
		"/api/profiles/targetuser",
		nil,
	)
	ctx.Request = httpReq
	ctx.Set("authenticated", false)

	r.ServeHTTP(w, httpReq)

	// Then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{
		"profile": {
			"username": "targetuser",
			"bio": "target bio",
			"image": "",
			"following": false
		}
	}`, w.Body.String())
}
