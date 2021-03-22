package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/layouts/db/sqlc"
)

type GetLayoutsRequest struct {
	AreaMin         float32 `form:"area_min"`
	AreaMax         float32 `form:"area_max"`
	LivingAreaMin   float32 `form:"living_area_min"`
	LivingAreaMax   float32 `form:"living_area_max"`
	CitchenAreaMin  float32 `form:"citching_area_min"`
	CitchenAreaMax  float32 `form:"citching_area_max"`
	CitchenAreaDesc bool    `form:"citchen_area_desc"`
	CitchenAreaAsc  bool    `form:"citchen_area_asc"`
	LivingAreaDesc  bool    `form:"living_area_desc"`
	LivingAreaAsc   bool    `form:"living_area_asc"`
	AreaDesc        bool    `form:"area_desc"`
	AreaAsc         bool    `form:"area_asc"`
	OffSet          float32 `form:"off_set"`
	Parent          int64   `form:"parent"`
}

type GetLayoutRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

func (server *Server) GetLayoutsList(ctx *gin.Context) {
	var req GetLayoutsRequest
	err := ctx.ShouldBindQuery(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetFilteredLayoutsParams{
		OffSet:          1,
		Parent:          1,
		AreaMax:         6969,
		LivingAreaMax:   6969,
		CitchenAreaMax:  6969,
		CitchenAreaDesc: req.CitchenAreaDesc,
		CitchenAreaAsc:  req.CitchenAreaAsc,
		LivingAreaDesc:  req.LivingAreaDesc,
		LivingAreaAsc:   req.LivingAreaAsc,
		AreaDesc:        req.AreaDesc,
		AreaAsc:         req.AreaAsc,
	}

	if req.Parent != 0 {
		arg.Parent = int32(req.Parent)
	}
	if req.OffSet != 0 {
		arg.OffSet = int32(req.OffSet)
	}

	if req.AreaMin >= 0 && req.AreaMax > 0 && req.AreaMin < req.AreaMax {
		arg.AreaMin = float64(req.AreaMin)
		arg.AreaMax = float64(req.AreaMax)
	}

	if req.LivingAreaMin >= 0 && req.LivingAreaMax > 0 && req.LivingAreaMin < req.LivingAreaMax {
		arg.LivingAreaMin = float64(req.LivingAreaMin)
		arg.LivingAreaMax = float64(req.LivingAreaMax)
	}
	if req.CitchenAreaMin >= 0 && req.CitchenAreaMax > 0 && req.CitchenAreaMin < req.CitchenAreaMax {
		arg.CitchenAreaMin = float64(req.CitchenAreaMin)
		arg.CitchenAreaMax = float64(req.CitchenAreaMax)
	}

	layouts, err := server.store.GetFilteredLayouts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, layouts)
	return
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
