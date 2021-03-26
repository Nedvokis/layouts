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

type Layout struct {
	ID          int64   `json:"id"`
	Parent      int64   `json:"parent"`
	Area        float64 `json:"area"`
	CitchenArea float64 `json:"citchen_area"`
	Door        int32   `json:"door"`
	Floor       int32   `json:"floor"`
	BitrixID    int32   `json:"bitrix_id"`
	LayoutID    int32   `json:"layout_id"`
	LivingArea  float64 `json:"living_area"`
	Num         string  `json:"num"`
	Price       int32   `json:"price"`
	Status      int32   `json:"status"`
	Type        int32   `json:"type"`
	Room        int32   `json:"room"`
	LayoutsUrl  string  `json:"layouts_url"`
	SvgPath     string  `json:"svg_path"`
}

type GetLayoutRequest struct {
	ID int64 `uri:"id" binding:"required,min=1`
}

type LayoutData struct {
	Layouts []Layout `json:"layouts"`
	Length  int      `json:"length"`
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

	refactoredAllLayouts := []Layout{}
	for _, s := range allLayouts {
		refactoredAllLayouts = append(refactoredAllLayouts, Layout{
			ID:          s.ID,
			Parent:      s.Parent,
			Area:        s.Area.Float64,
			CitchenArea: s.CitchenArea.Float64,
			Door:        s.Door.Int32,
			Floor:       s.Floor.Int32,
			BitrixID:    s.BitrixID.Int32,
			LayoutID:    s.LayoutID.Int32,
			LivingArea:  s.LivingArea.Float64,
			Num:         s.Num.String,
			Price:       s.Price.Int32,
			Status:      s.Status.Int32,
			Type:        s.Type.Int32,
			Room:        s.Room.Int32,
			LayoutsUrl:  s.LayoutsUrl.String,
			SvgPath:     s.SvgPath.String,
		})
	}

	if req.GetAll == true {
		ctx.JSON(http.StatusOK, LayoutData{
			Layouts: refactoredAllLayouts,
			Length:  length,
		})
		return
	}

	layouts, err := server.store.GetFilteredLayouts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refactoredLayouts := []Layout{}
	for _, s := range layouts {
		refactoredLayouts = append(refactoredLayouts, Layout{
			ID:          s.ID,
			Parent:      s.Parent,
			Area:        s.Area.Float64,
			CitchenArea: s.CitchenArea.Float64,
			Door:        s.Door.Int32,
			Floor:       s.Floor.Int32,
			BitrixID:    s.BitrixID.Int32,
			LayoutID:    s.LayoutID.Int32,
			LivingArea:  s.LivingArea.Float64,
			Num:         s.Num.String,
			Price:       s.Price.Int32,
			Status:      s.Status.Int32,
			Type:        s.Type.Int32,
			Room:        s.Room.Int32,
			LayoutsUrl:  s.LayoutsUrl.String,
			SvgPath:     s.SvgPath.String,
		})
	}

	data := LayoutData{
		Layouts: refactoredLayouts,
		Length:  length,
	}

	ctx.JSON(http.StatusOK, data)
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
	data := Layout{
		ID:          layout.ID,
		Parent:      layout.Parent,
		Area:        layout.Area.Float64,
		CitchenArea: layout.CitchenArea.Float64,
		Door:        layout.Door.Int32,
		Floor:       layout.Floor.Int32,
		BitrixID:    layout.BitrixID.Int32,
		LayoutID:    layout.LayoutID.Int32,
		LivingArea:  layout.LivingArea.Float64,
		Num:         layout.Num.String,
		Price:       layout.Price.Int32,
		Status:      layout.Status.Int32,
		Type:        layout.Type.Int32,
		Room:        layout.Room.Int32,
		LayoutsUrl:  layout.LayoutsUrl.String,
		SvgPath:     layout.SvgPath.String,
	}

	ctx.JSON(http.StatusOK, data)
}
