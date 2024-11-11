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
	"os"
	"testing"
)

var r *gin.Engine
var g *gin.RouterGroup
var w *httptest.ResponseRecorder
var mockLogic *app.MockLogic

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	g = r.Group("/api")
	w = httptest.NewRecorder()

	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()
	mockLogic = app.NewMockLogic(ctrl)

	code := m.Run()

	os.Exit(code)
}

func TestRoutes_Post_Users_Login_Success(t *testing.T) {
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
