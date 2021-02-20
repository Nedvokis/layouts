package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/layouts/db/sqlc"
)

type createComplexRequest struct {
	BitrixID int64  `json:"bitrix_id"`
	Name     string `json:"name"`
}

func (server *Server) createComplex(ctx *gin.Context) {
	var req createComplexRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateComplexParams{
		BitrixID: req.BitrixID,
		Name:     sql.NullString{String: req.Name, Valid: true},
	}
	complex, err := server.store.CreateComplex(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, complex)

}
