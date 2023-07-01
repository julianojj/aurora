package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type ProjectRoute struct {
	r                       *gin.Engine
	createProjectController *controllers.CreateProjectController
	getProjectController    *controllers.GetProjectController
}

func NewProjectRoute(
	r *gin.Engine,
	createProjectController *controllers.CreateProjectController,
	getProjectController *controllers.GetProjectController,
) *ProjectRoute {
	return &ProjectRoute{
		r,
		createProjectController,
		getProjectController,
	}
}

func (pr *ProjectRoute) Register() {
	pr.r.POST("/create_project", pr.createProjectController.Handle)
	pr.r.GET("/get_project/:id", pr.getProjectController.Handle)
}
