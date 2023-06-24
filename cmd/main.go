package main

import (
	"net/http"
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
	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	app.Use(cors.New(config))

	// Repositories
	fileRepository := repository.NewFileRepositoryMemory()
	projectRepository := repository.NewProjectRepositoryMemory()

	// Adapters
	bucket := adapters.NewMinio(os.Getenv("BUCKET_NAME"))
	err := bucket.CreateBucket()
	if err != nil {
		panic(err)
	}
	// Usecases
	createProject := usecases.NewCreateProject(projectRepository)
	getProject := usecases.NewGetProject(projectRepository)
	uploadFile := usecases.NewUploadFile(fileRepository, bucket)
	removeFile := usecases.NewRemoveFile(fileRepository)
	getUploads := usecases.NewGetUploads(fileRepository)

	// Controllers
	createProjectController := controllers.NewCreateProjectController(createProject)
	getProjectController := controllers.NewGetProjectController(getProject)
	getUploadsController := controllers.NewGetUploadsController(getUploads)
	uploadFileController := controllers.NewUploadFileController(uploadFile)
	removeFileController := controllers.NewRemoveFileController(removeFile)

	// Routes
	routes.NewUploadRoute(
		app,
		uploadFileController,
		getUploadsController,
		removeFileController,
	).Register()
	routes.NewProjectRoute(
		app,
		createProjectController,
		getProjectController,
	).Register()

	app.GET("/assets/:id", func(c *gin.Context) {
		assetID := c.Param("id")
		asset, err := bucket.GetObject(assetID)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"code":    http.StatusBadRequest,
			})
		}
		c.Writer.Write(asset)
	})
	err = app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
