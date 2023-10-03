package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julianojj/aurora/internal/core/usecases"
	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
	"github.com/julianojj/aurora/internal/infra/api/middlewares"
	"github.com/julianojj/aurora/internal/infra/api/routes"
	"github.com/julianojj/aurora/internal/infra/repository"
	"go.uber.org/zap"
)

func main() {
	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	app.Use(cors.New(config))
	logger, _ := zap.NewProduction()

	// Repositories
	fileRepository := repository.NewFileRepositoryMemory()
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()

	// Adapters
	bucket := adapters.NewMinio(os.Getenv("BUCKET_NAME"))
	err := bucket.CreateBucket()
	if err != nil {
		panic(err)
	}
	// Usecases
	createProject := usecases.NewCreateProject(projectRepository, logger)
	getProject := usecases.NewGetProject(projectRepository)
	createArtboard := usecases.NewCreateArtboard(projectRepository, artboardRepository, logger)
	uploadFile := usecases.NewUploadFile(fileRepository, bucket)
	removeFile := usecases.NewRemoveFile(fileRepository, bucket, logger)
	getUploads := usecases.NewGetUploads(fileRepository)
	getAsset := usecases.NewGetAsset(bucket, logger)

	// Controllers
	createProjectController := controllers.NewCreateProjectController(createProject)
	getProjectController := controllers.NewGetProjectController(getProject)
	createArtboardController := controllers.NewCreateArtboardController(createArtboard)
	getUploadsController := controllers.NewGetUploadsController(getUploads)
	uploadFileController := controllers.NewUploadFileController(uploadFile)
	removeFileController := controllers.NewRemoveFileController(removeFile)
	getAssetController := controllers.NewGetAssetController(getAsset)

	// Middlewares
	app.Use(middlewares.ErrorHandler())

	// Routes
	routes.NewUploadRoute(
		app,
		uploadFileController,
		getUploadsController,
		removeFileController,
		getAssetController,
	).Register()
	routes.NewProjectRoute(
		app,
		createProjectController,
		getProjectController,
		createArtboardController,
	).Register()

	err = app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
