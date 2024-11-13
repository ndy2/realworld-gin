package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	authapp "ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/user/app"
	"ndy/realworld-gin/internal/user/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes_Post_Users_Success(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

	// Given
	data, _ := json.Marshal(map[string]dto.RegistrationRequest{
		"user": {
			Username: "jakejake",
			Email:    "jake@jake.jake",
			Password: "Passw0rd!",
		},
	})

	mockLogic.EXPECT().Register("jakejake", "jake@jake.jake", "Passw0rd!").Return(1, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"POST",
		"/api/users",
		bytes.NewReader(data),
	)
	r.ServeHTTP(w, httpReq)

	// Then
	wantStatus := http.StatusOK
	wantResp, _ := json.Marshal(map[string]dto.RegistrationResponse{"user": {
		Email:    "jake@jake.jake",
		Username: "jakejake",
		Token:    "",
		Bio:      "",
		Image:    "",
	}})

	assert.Equal(t, wantStatus, w.Code)
	assert.JSONEq(t, string(wantResp), w.Body.String())
}

func TestRoutes_Get_User_Success(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

	// Given
	resp := dto.GetCurrentUserResponse{
		Email:    "jake@jake.jake",
		Token:    "",
		Username: "jakejake",
		Bio:      "jake bio",
		Image:    "http://image.jake.image",
	}
	mockLogic.EXPECT().GetCurrentUser(1, 1).Return(resp, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"GET",
		"/api/user",
		nil,
	)

	testToken, _ := authapp.Generate(1, 1, "jakejake")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Token %s", testToken))

	r.ServeHTTP(w, httpReq)

	// Then
	wantStatus := http.StatusOK
	resp.Token = testToken
	wantResp, _ := json.Marshal(map[string]dto.GetCurrentUserResponse{"user": resp})

	assert.Equal(t, wantStatus, w.Code)
	assert.JSONEq(t, string(wantResp), w.Body.String())
}

func TestRoutes_Put_User_Success(t *testing.T) {
	r, g, w, ctrl, mockLogic := setupRoutes(t)
	defer ctrl.Finish()

	// Given
	data, _ := json.Marshal(map[string]dto.UpdateUserRequest{
		"user": {
			Email:    "jake@jake.jake",
			Username: "jakejake",
			Password: "Passw0rd!",
			Bio:      "jake bio",
			Image:    "http://image.jake.image",
		},
	})

	resp := dto.UpdateUserResponse{
		Email:    "jake@jake.jake",
		Username: "jakejake",
		Token:    "",
		Bio:      "jake bio",
		Image:    "http://image.jake.image",
	}

	mockLogic.EXPECT().UpdateUser(gomock.Any(), "jake@jake.jake", "jakejake", "Passw0rd!", "http://image.jake.image", "jake bio").Return(resp, nil)
	var logic app.Logic = mockLogic

	// When
	Routes(g, &logic)
	httpReq, _ := http.NewRequest(
		"PUT",
		"/api/user",
		bytes.NewReader(data),
	)

	testToken, _ := authapp.Generate(1, 1, "jakejake")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Token %s", testToken))

	r.ServeHTTP(w, httpReq)

	// Then
	wantStatus := http.StatusOK
	resp.Token = testToken
	wantResp, _ := json.Marshal(map[string]dto.UpdateUserResponse{"user": resp})

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
