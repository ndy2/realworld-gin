package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	authapp "ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/profile/app"
	"ndy/realworld-gin/internal/profile/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes_Get_Profiles_Username_Success_Authenticated(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

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
	testToken, _ := authapp.Generate(1, 1, "testuser")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Token %s", testToken))

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

func TestRoutes_Get_Profiles_Username_Success_Unauthenticated(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

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

func setupRoutes(t *testing.T) (*gin.Engine, *gin.RouterGroup, *httptest.ResponseRecorder, *gomock.Controller, *app.MockLogic) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	g := r.Group("/api")
	w := httptest.NewRecorder()
	ctrl := gomock.NewController(t)
	mockLogic := app.NewMockLogic(ctrl)

	return r, g, w, ctrl, mockLogic
}
