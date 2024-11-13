package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes_Post_Users_Login_Success(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

	// Given
	data, _ := json.Marshal(map[string]dto.LoginRequest{
		"user": {
			Email:    "jake@jake.jake",
			Password: "jakejake",
		},
	})
	resp := dto.LoginResponse{
		Email:    "jake@jake.jake",
		Token:    "jwt.token.here",
		Username: "jake",
		Bio:      "I work at statefarm",
		Image:    "",
	}
	mockLogic.EXPECT().Login("jake@jake.jake", "jakejake").Return(resp, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"POST",
		"/api/users/login",
		bytes.NewReader(data),
	)
	r.ServeHTTP(w, httpReq)

	// Then
	wantStatus := http.StatusOK
	wantResp, _ := json.Marshal(map[string]dto.LoginResponse{"user": resp})

	assert.Equal(t, wantStatus, w.Code)
	assert.JSONEq(t, string(wantResp), w.Body.String())
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
