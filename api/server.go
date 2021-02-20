package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/layouts/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	// router.POST("/complex", store.createComplex)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific addres.
func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
