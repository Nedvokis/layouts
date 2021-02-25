package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetLayoutsList(ctx *gin.Context) {
	layouts, err := server.store.GetAllListLayouts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, layouts)
}
