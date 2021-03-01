package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetComplexesList(ctx *gin.Context) {
	complexes, err := server.store.GetListAllComplexes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, complexes)
}

type GetComplexRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

func (server *Server) GetComplex(ctx *gin.Context) {
	var req GetComplexRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	complex, err := server.store.GetComplex(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, complex)
}
