package v0

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
}

var _v0 = &service{}

func FailedJSON(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "failed",
		"Message": msg,
	})
}

func Register(group *gin.RouterGroup) {
	_v0.Hash(group)
}

func (service) Hash(group *gin.RouterGroup) {
	group.GET(":hash", func(ctx *gin.Context) {
		hash := ctx.Param("hash")
		if hash == "" {
			FailedJSON(ctx, fmt.Sprintf("%v not found", hash))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "success",
			"Data":   hash,
		})
	})
}
