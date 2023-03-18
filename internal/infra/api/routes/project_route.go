package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type ProjectRoute struct {
	r                       *gin.Engine
	createProjectController *controllers.CreateProjectController
}

func NewProjectRoute(
	r *gin.Engine,
	createProjectController *controllers.CreateProjectController,
) *ProjectRoute {
	return &ProjectRoute{
		r,
		createProjectController,
	}
}

func (pr *ProjectRoute) Register() {
	pr.r.POST("/projects", pr.createProjectController.Handle)
}
