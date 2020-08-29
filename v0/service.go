package v0

import (
	"github.com/glvd/link-rest/model"
	"github.com/xormsharp/xorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	db *xorm.Engine
}

var _v0 = &service{}

func FailedJSON(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "failed",
		"Message": msg,
	})
}

func Register(db *xorm.Engine, group *gin.RouterGroup) {
	_v0.db = db
	_v0.total(group)
}

func (s service) total(group *gin.RouterGroup) {
	group.GET("/all", func(ctx *gin.Context) {
		page := model.NewPage(new([]model.Media))
		page.Parse(ctx.Request.URL.Query())

		find, err := page.Find(s.db.NewSession())
		if err != nil {
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Status": "success",
			"Data":   find,
		})
	})
}
