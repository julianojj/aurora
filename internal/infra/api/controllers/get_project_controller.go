package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type GetProjectController struct {
	getProject *usecases.GetProject
}

func NewGetProjectController(getProject *usecases.GetProject) *GetProjectController {
	return &GetProjectController{
		getProject,
	}
}

func (gpc *GetProjectController) Handle(c *gin.Context) {
	project, err := gpc.getProject.Execute(c.Param("id"))
	if err == nil {
		c.JSON(http.StatusOK, project)
		return
	}
	c.Error(err)
}
