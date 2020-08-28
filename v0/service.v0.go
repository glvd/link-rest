package v0

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FailedJSON(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "failed",
		"Message": msg,
	})
}

func RegisterV0(group *gin.RouterGroup) {
	group.GET(":hash", func(ctx *gin.Context) {
		hash := ctx.Param("hash")
		if hash == "" {
			return
		}
	})
}
