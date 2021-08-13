package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLittersRequest struct {
	Parent int64 `form:"parent" binding:"min=1"`
}

func (server *Server) GetLittersList(ctx *gin.Context) {
	var req GetLittersRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		litters, err := server.store.GetListAllLitters(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, litters)
		return
	}

	litters, err := server.store.GetListLittersByParent(ctx, req.Parent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, litters)
	return

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

type GetLitterByBxID struct {
	BitrixID int64 `form:"bitrix_id" json:"bitrix_id"`
}

func (server *Server) GetLitterByBxID(ctx *gin.Context) {
	var req GetLitterByBxID

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	litter, err := server.store.GetLitterByBxID(ctx, req.BitrixID)
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
