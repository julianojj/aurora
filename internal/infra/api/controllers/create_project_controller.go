package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type CreateProjectController struct {
	createProject *usecases.CreateProject
}

func NewCreateProjectController(createProject *usecases.CreateProject) *CreateProjectController {
	return &CreateProjectController{
		createProject,
	}
}

func (cpc *CreateProjectController) Handle(c *gin.Context) {
	var input usecases.CreateProjectInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}
	project, err := cpc.createProject.Execute(input)
	if err == nil {
		c.JSON(http.StatusCreated, project)
		return
	}
	c.Error(err)
}
