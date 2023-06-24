package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type UploadRoute struct {
	r                    *gin.Engine
	uploadFileController *controllers.UploadFileController
	getUploadsController *controllers.GetUploadsController
	removeFileController controllers.RemoveFileController
}

func NewUploadRoute(
	r *gin.Engine,
	uploadFileController *controllers.UploadFileController,
	getUploadsController *controllers.GetUploadsController,
	removeFileCOntroller *controllers.RemoveFileController,
) *UploadRoute {
	return &UploadRoute{
		r,
		uploadFileController,
		getUploadsController,
		*removeFileCOntroller,
	}
}

func (u *UploadRoute) Register() {
	u.r.POST("/upload", u.uploadFileController.Handle)
	u.r.GET("/uploads", u.getUploadsController.Handle)
	u.r.DELETE("/uploads/:id", u.removeFileController.Handle)
}
