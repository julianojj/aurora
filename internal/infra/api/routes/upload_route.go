package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/infra/api/controllers"
)

type UploadRoute struct {
	r                    *gin.Engine
	uploadFileController *controllers.UploadFileController
}

func NewUploadFileRoute(
	r *gin.Engine,
	uploadFileController *controllers.UploadFileController,
) *UploadRoute {
	return &UploadRoute{
		r,
		uploadFileController,
	}
}

func (u *UploadRoute) Register() {
	u.r.POST("/upload", u.uploadFileController.Handle)
}
