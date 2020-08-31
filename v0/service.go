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
		page := model.NewPage(new([]model.Media))
		page.Parse(ctx.Request.URL.Query())

		s := s.db.NewSession()
		defer s.Close()
		s.Table(new(model.Media))
		find, err := page.Find(s)
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Status": "success",
			"Data":   find,
		})
	})
}