package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type CreateArtboardController struct {
	createArtboard *usecases.CreateArtboard
}

func NewCreateArtboardController(createArtboard *usecases.CreateArtboard) *CreateArtboardController {
	return &CreateArtboardController{
		createArtboard,
	}
}

func (cac *CreateArtboardController) Handle(c *gin.Context) {
	var input usecases.CreateArtboardInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}
	err = cac.createArtboard.Execute(input)
	if err == nil {
		c.JSON(201, nil)
		return
	}
	c.Error(err)
}
