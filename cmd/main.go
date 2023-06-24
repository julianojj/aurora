package main

import (
	"os"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	fileRepository := repository.NewFileRepositoryMemory()
	projectRepository := repository.NewProjectRepositoryMemory()
	bucket := adapters.NewMinio(os.Getenv("BUCKET_NAME"))
	err := bucket.CreateBucket()
	if err != nil {
		panic(err)
	}
	createProject := usecases.NewCreateProject(projectRepository)
	createProjectController := controllers.NewCreateProjectController(createProject)
	getProject := usecases.NewGetProject(projectRepository)
	getProjectController := controllers.NewGetProjectController(getProject)
	uploadFile := usecases.NewUploadFile(fileRepository, bucket)
	uploadFileController := controllers.NewUploadFileController(uploadFile)
	removeFile := usecases.NewRemoveFile(fileRepository)
	removeFileController := controllers.NewRemoveFileController(removeFile)
	getUploads := usecases.NewGetUploads(fileRepository)
	getUploadsController := controllers.NewGetUploadsController(getUploads)
	routes.NewUploadRoute(
		r,
		uploadFileController,
		getUploadsController,
		removeFileController,
	).Register()
	routes.NewProjectRoute(
		r,
		createProjectController,
		getProjectController,
	).Register()
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
