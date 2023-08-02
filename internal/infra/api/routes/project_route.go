package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type ProjectRoute struct {
	r                        *gin.Engine
	createProjectController  *controllers.CreateProjectController
	getProjectController     *controllers.GetProjectController
	createArtboardController *controllers.CreateArtboardController
}

func NewProjectRoute(
	r *gin.Engine,
	createProjectController *controllers.CreateProjectController,
	getProjectController *controllers.GetProjectController,
	createArtboardController *controllers.CreateArtboardController,
) *ProjectRoute {
	return &ProjectRoute{
		r,
		createProjectController,
		getProjectController,
		createArtboardController,
	}
}

func (pr *ProjectRoute) Register() {
	pr.r.POST("/create_project", pr.createProjectController.Handle)
	pr.r.GET("/get_project/:id", pr.getProjectController.Handle)
	pr.r.POST("/create_artboard", pr.createArtboardController.Handle)
}
