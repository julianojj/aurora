package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julianojj/aurora/internal/core/usecases"
	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
	"github.com/julianojj/aurora/internal/infra/api/routes"
	"github.com/julianojj/aurora/internal/infra/repository"
)

func main() {
	r := gin.Default()
	fileRepository := repository.NewFileRepositoryMemory()
	projectRepository := repository.NewProjectRepositoryMemory()
	bucket := adapters.NewMinio()
	err := bucket.CreateBucket()
	if err != nil {
		panic(err)
	}
	createProject := usecases.NewCreateProject(projectRepository)
	createProjectController := controllers.NewCreateProjectController(createProject)
	uploadFile := usecases.NewUploadFile(fileRepository, bucket)
	uploadFileController := controllers.NewUploadFileController(uploadFile)
	getUploads := usecases.NewGetUploads(fileRepository)
	getUploadsController := controllers.NewGetUploadsController(getUploads)
	routes.NewUploadRoute(
		r,
		uploadFileController,
		getUploadsController,
	).Register()
	routes.NewProjectRoute(
		r,
		createProjectController,
	).Register()
	r.Run(":8080")
}
