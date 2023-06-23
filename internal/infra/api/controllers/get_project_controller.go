package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/exceptions"
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
	switch err.(type) {
	case *exceptions.NotFoundException:
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"code":    http.StatusNotFound,
		})
	case *exceptions.ValidationException:
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
			"code":    http.StatusUnprocessableEntity,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"code":    http.StatusInternalServerError,
		})
	}
}
