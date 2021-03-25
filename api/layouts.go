package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/layouts/db/sqlc"
)

const (
	MAX_VALUE = 99999
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
	Room            int64   `form:"room"`
	GetAll          bool    `form:"get_all"`
}

type GetLayoutRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

type LayoutData struct {
	Layouts []db.Layout `json:"layouts"`
	Length  int         `json:"length"`
}

func (server *Server) GetLayoutsList(ctx *gin.Context) {
	var req GetLayoutsRequest
	err := ctx.ShouldBindQuery(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetFilteredLayoutsParams{
		AreaMax:         MAX_VALUE,
		LivingAreaMax:   MAX_VALUE,
		CitchenAreaMax:  MAX_VALUE,
		OffSet:          int32(req.OffSet),
		Room:            int32(req.Room),
		Parent:          int32(req.Parent),
		CitchenAreaDesc: req.CitchenAreaDesc,
		CitchenAreaAsc:  req.CitchenAreaAsc,
		LivingAreaDesc:  req.LivingAreaDesc,
		LivingAreaAsc:   req.LivingAreaAsc,
		AreaDesc:        req.AreaDesc,
		AreaAsc:         req.AreaAsc,
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

	argLength := db.GetFilteredLayoutsLengthParams{
		AreaMax:         MAX_VALUE,
		LivingAreaMax:   MAX_VALUE,
		CitchenAreaMax:  MAX_VALUE,
		Room:            int32(req.Room),
		Parent:          int32(req.Parent),
		CitchenAreaDesc: req.CitchenAreaDesc,
		CitchenAreaAsc:  req.CitchenAreaAsc,
		LivingAreaDesc:  req.LivingAreaDesc,
		LivingAreaAsc:   req.LivingAreaAsc,
		AreaDesc:        req.AreaDesc,
		AreaAsc:         req.AreaAsc,
	}
	if req.AreaMin >= 0 && req.AreaMax > 0 && req.AreaMin < req.AreaMax {
		argLength.AreaMin = float64(req.AreaMin)
		argLength.AreaMax = float64(req.AreaMax)
	}

	if req.LivingAreaMin >= 0 && req.LivingAreaMax > 0 && req.LivingAreaMin < req.LivingAreaMax {
		argLength.LivingAreaMin = float64(req.LivingAreaMin)
		argLength.LivingAreaMax = float64(req.LivingAreaMax)
	}

	if req.CitchenAreaMin >= 0 && req.CitchenAreaMax > 0 && req.CitchenAreaMin < req.CitchenAreaMax {
		argLength.CitchenAreaMin = float64(req.CitchenAreaMin)
		argLength.CitchenAreaMax = float64(req.CitchenAreaMax)
	}

	allLayouts, err := server.store.GetFilteredLayoutsLength(ctx, argLength)
	length := len(allLayouts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.GetAll == true {
		ctx.JSON(http.StatusOK, LayoutData{
			Layouts: allLayouts,
			Length:  length,
		})
		return
	}

	layouts, err := server.store.GetFilteredLayouts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, LayoutData{
		Layouts: layouts,
		Length:  length,
	})
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
