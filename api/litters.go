package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetLittersList(ctx *gin.Context) {
	litters, err := server.store.GetListAllLitters(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, litters)
}

type GetLitterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

func (server *Server) GetLitter(ctx *gin.Context) {
	var req GetLitterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	litter, err := server.store.GetLitter(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, litter)
}
