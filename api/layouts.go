package api

import (
	"database/sql"
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

type GetLayoutRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

func (server *Server) GetLayout(ctx *gin.Context) {
	var req GetLayoutRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	layout, err := server.store.GetLayout(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, layout)
}
