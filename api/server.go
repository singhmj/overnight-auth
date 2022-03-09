package api

import (
	"auth/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	server.router.POST("/register", server.createUser)
	server.router.GET("/login", server.loginUser)
	server.router.GET("/verify", server.verifyUser)
	server.router.GET("/updatePassword", server.updateUserPassword)
	server.router.PUT("/updateStatus", server.updateUserStatus)

	server.router.POST("/profiles", server.createProfile)
	server.router.GET("/profiles/:id", server.getProfile)
	server.router.GET("/profiles", server.listProfiles)
	server.router.PUT("/profiles", server.updateProfile)
	// server.router.DELETE("/profiles", server.listProfiles)

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
