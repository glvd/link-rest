package v0

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glvd/link-rest/model"
	"github.com/goextension/log"
	"github.com/xormsharp/xorm"
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
	group.GET("/show", func(ctx *gin.Context) {
		page := model.Page(new([]model.Media), ctx.Request)

		sess := s.db.Table(new(model.Media)).Join("left join", "file", "file_id = file.id").Join("left join", "info", "info_id = info.id")
		find, err := page.Find(sess)
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	})
}
