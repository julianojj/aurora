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
	routes.NewUploadFileRoute(
		r,
		uploadFileController,
	).Register()
	routes.NewProjectRoute(
		r,
		createProjectController,
	).Register()
	r.GET("/uploads", func(c *gin.Context) {
		files, _ := fileRepository.FindAll()
		c.JSON(200, files)
	})
	r.Run(":8080")
}
