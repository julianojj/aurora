package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type GetArtboardsController struct {
	getArtboards *usecases.GetArtboards
}

func NewGetArtboardsController(getArtboards *usecases.GetArtboards) *GetArtboardsController {
	return &GetArtboardsController{
		getArtboards,
	}
}

func (gac *GetArtboardsController) Handle(ctx *gin.Context) {
	projectId := ctx.Param("id")
	output, err := gac.getArtboards.Execute(projectId)
	if err == nil {
		ctx.JSON(200, output)
		return
	}
	ctx.Error(err)
}
