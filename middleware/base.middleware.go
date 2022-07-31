package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/response"
	"gitlab.com/mr687/privy-be-test-go/service"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Content-Type", "application/json")
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Status:     false,
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       []string{err.Error()},
			})
		}

		c.Next()
	}
}
