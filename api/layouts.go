package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLayoutsRequest struct {
	Parent int64 `form:"parent" binding:"min=1"`
}

func (server *Server) GetLayoutsList(ctx *gin.Context) {
	var req GetLayoutsRequest
	err := ctx.BindQuery(&req)
	if err == nil {
		layouts, err := server.store.GetLayoutByLitter(ctx, req.Parent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, layouts)
		return
	}

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
