package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type GetUploadsController struct {
	getUploads *usecases.GetUploads
}

func NewGetUploadsController(
	getUploads *usecases.GetUploads,
) *GetUploadsController {
	return &GetUploadsController{
		getUploads,
	}
}

func (guc *GetUploadsController) Handle(c *gin.Context) {
	uploads, err := guc.getUploads.Execute()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uploads)
}
