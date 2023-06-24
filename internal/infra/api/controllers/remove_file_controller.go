package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/usecases"
)

type RemoveFileController struct {
	removeFile *usecases.RemoveFile
}

func NewRemoveFileController(
	removeFile *usecases.RemoveFile,
) *RemoveFileController {
	return &RemoveFileController{
		removeFile,
	}
}

func (rfc *RemoveFileController) Handle(c *gin.Context) {
	err := rfc.removeFile.Execute(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusNoContent, map[string]any{})
}
