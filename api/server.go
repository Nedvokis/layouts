package api

import (
	"net/http"

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
	router.OPTIONS("/api2/complexes", preflight)
	router.GET("/api2/litters", server.GetLittersList)
	router.GET("/api2/litters/:id", server.GetLitter)
	router.OPTIONS("/api2/litters", preflight)
	router.GET("/api2/layouts", server.GetLayoutsList)
	router.GET("/api2/layouts/:id", server.GetLayout)
	router.OPTIONS("/api2/layouts", preflight)
	router.GET("/api2/setDb", server.SetDb)
	router.GET("/api2/setSvg", server.SetSvg)

	server.router = router
	return server
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

// Start runs the HTTP server on a specific addres.
func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
