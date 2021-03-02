package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/layouts/db/sqlc"
)

type GetLayoutsRequest struct {
	Parent int64 `form:"parent" binding:"min=1,gtfield=Door"`
	Door   int32 `form:"door" binding:"min=1"`
}

func (server *Server) GetLayoutsList(ctx *gin.Context) {
	var req GetLayoutsRequest
	err := ctx.BindQuery(&req)
	if err == nil {
		arr := db.GetLayoutByLitterParams{
			Parent: req.Parent,
			Door: sql.NullInt32{
				Int32: req.Door,
				Valid: true,
			},
		}
		layouts, err := server.store.GetLayoutByLitter(ctx, arr)
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
