package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type UploadRoute struct {
	r                    *gin.Engine
	uploadFileController *controllers.UploadFileController
	getUploadsController *controllers.GetUploadsController
}

func NewUploadRoute(
	r *gin.Engine,
	uploadFileController *controllers.UploadFileController,
	getUploadsController *controllers.GetUploadsController,
) *UploadRoute {
	return &UploadRoute{
		r,
		uploadFileController,
		getUploadsController,
	}
}

func (u *UploadRoute) Register() {
	u.r.POST("/upload", u.uploadFileController.Handle)
	u.r.GET("/uploads", u.getUploadsController.Handle)
}
