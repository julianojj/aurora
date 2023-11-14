package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/aurora/internal/core/exceptions"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, err := range ctx.Errors {
			switch err.Err.(type) {
			case *exceptions.NotFoundException:
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
					"code":    http.StatusNotFound,
				})
			case *exceptions.ValidationException:
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": err.Error(),
					"code":    http.StatusUnprocessableEntity,
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
					"code":    http.StatusInternalServerError,
				})
			}
		}
	}
}
