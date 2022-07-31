package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/response"
)

func (s *Server) Index(c *gin.Context) {
	response.Ok(c, "Simple API - Privy BE Test", nil)
}
