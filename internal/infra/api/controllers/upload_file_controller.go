package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type UploadFileController struct {
	uploadFile *usecases.UploadFile
}

func NewUploadFileController(
	uploadFile *usecases.UploadFile,
) *UploadFileController {
	return &UploadFileController{
		uploadFile,
	}
}

func (u *UploadFileController) Handle(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error uploading file",
			"code":    400,
		})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error to open file headers",
			"code":    400,
		})
		return
	}
	input := usecases.UploadFileInput{
		Name:     fileHeader.Filename,
		Size:     fileHeader.Size,
		Mimetype: fileHeader.Header.Get("Content-Type"),
		Reader:   file,
	}
	err = u.uploadFile.Execute(input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    500,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success upload file",
		"code":    200,
	})
}
