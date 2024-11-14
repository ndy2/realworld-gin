package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"ndy/realworld-gin/internal/auth/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOptionalAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	token, _ := app.Generate(1, 1, "testuser")

	tests := []struct {
		name           string
		authHeader     string
		expectedAuth   bool
		expectedStatus int
	}{
		{"No Authorization Header", "", false, http.StatusOK},
		{"With Invalid Authorization Header", "Token invalid_valid_token", true, http.StatusUnauthorized},
		{"With Valid Authorization Header", fmt.Sprintf("Token %s", token), true, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a response recorder
			w := httptest.NewRecorder()

			// Set up a test Gin context with a request
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			c.Request = req

			// Initialize middleware and execute it
			middleware := OptionalAuth()
			middleware(c)

			// Check if the `authenticated` value in the context is set as expected
			authenticated, exists := c.Get("authenticated")
			assert.True(t, exists)
			assert.Equal(t, tt.expectedAuth, authenticated)

			// Check response status code
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
