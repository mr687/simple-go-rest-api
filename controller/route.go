package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/response"
)

func (s *Server) InitializeRoutes() {
	// Index Route
	s.Router.GET("/", s.Index)

	// API Routes
	api := s.Router.Group("/api/v1")
	{
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.Register)
			auth.POST("/login", s.Login)
			auth.POST("/logout", s.Logout)
		}

		// User Balance Routes
	}

	// Otherwise, just show Not Found 404
	s.Router.NoRoute(func(c *gin.Context) {
		response.NotFound(c)
	})

}
