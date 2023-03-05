package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julianojj/aurora/internal/core/usecases"
	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/repository"
)

func main() {
	r := gin.Default()
	fileRepository := repository.NewFileRepositoryMemory()
	bucket := adapters.NewMinio()
	err := bucket.CreateBucket()
	if err != nil {
		panic(err)
	}
	uploadFile := usecases.NewUploadFile(fileRepository, bucket)
	r.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error uploading file",
				"code":    400,
			})
		}
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error to open file headers",
				"code":    400,
			})
		}
		input := usecases.UploadFileInput{
			Name:     fileHeader.Filename,
			Size:     fileHeader.Size,
			Mimetype: fileHeader.Header.Get("Content-Type"),
			Reader:   file,
		}
		err = uploadFile.Execute(input)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    500,
			})
		}
		c.JSON(200, gin.H{
			"message": "Success upload file",
			"code":    200,
		})
	})
	r.GET("/uploads", func(c *gin.Context) {
		files, _ := fileRepository.FindAll()
		c.JSON(200, files)
	})
	r.Run(":8080")
}
