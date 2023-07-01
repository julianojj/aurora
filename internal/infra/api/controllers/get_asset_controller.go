package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type GetAssetController struct {
	getAsset *usecases.GetAsset
}

func NewGetAssetController(getAsset *usecases.GetAsset) *GetAssetController {
	return &GetAssetController{
		getAsset,
	}
}

func (gac *GetAssetController) Handle(c *gin.Context) {
	assetID := c.Param("id")
	asset, err := gac.getAsset.Execute(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}
	c.Writer.Write(asset)
}
