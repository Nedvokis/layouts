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
	router.GET("/api2/complexes", server.GetComplexesList)
	router.GET("/api2/complexes/:id", server.GetComplex)
	router.GET("/api2/litters", server.GetLittersList)
	router.GET("/api2/litters/:id", server.GetLitter)
	router.GET("/api2/layouts", server.GetLayoutsList)
	router.GET("/api2/layouts/:id", server.GetLayout)
	router.GET("/api2/setDb", server.SetDb)
	router.GET("/api2/setSvg", server.SetSvg)

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
