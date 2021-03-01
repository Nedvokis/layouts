package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	thirdparty "github.com/layouts/thirdParty"
)

type Success {
	status string
}

func (server *Server) SetDb(ctx *gin.Context) {
	thirdparty.GetLayouts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Success{
		status: "success"
	})
}
func (server *Server) SetSvg(ctx *gin.Context) {
	thirdparty.AddPathAndCreateSvgData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Success{
		status: "success"
	})
}

