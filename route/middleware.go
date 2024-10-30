package common

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

// HandleJsonRootMiddleware 는 JSON 의 root key 를 설정하고, root key 에 해당하는 데이터를 다음 핸들러로 전달하는 미들웨어를 생성합니다.
func HandleJsonRootMiddleware(reqJsonRootKey string, respJsonRootKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var root map[string]json.RawMessage

		// JSON 을 root map 으로 언마샬링
		if err := c.ShouldBindJSON(&root); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		// 지정된 reqJsonRootKey 에 해당하는 데이터를 찾아 설정
		if rootData, exists := root[reqJsonRootKey]; exists {
			c.Set("rootData", rootData)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data for key: " + reqJsonRootKey})
			c.Abort()
			return
		}

		// 다음 핸들러로 요청을 전달
		c.Next()

		// 감싸는 구조체 생성
		resp, _ := c.Get("resp")
		wrappedResponse := gin.H{respJsonRootKey: resp}
		c.IndentedJSON(http.StatusOK, wrappedResponse)
	}
}
