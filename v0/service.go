package v0

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/glvd/link-rest/model"
	"github.com/goextension/log"
	"gorm.io/gorm"
)

type service struct {
	db    *gorm.DB
	cache *persistence.InMemoryStore
}

var _v0 = &service{}

func FailedJSON(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "failed",
		"Message": msg,
	})
}

func Register(db *gorm.DB, group *gin.RouterGroup, cache *persistence.InMemoryStore) {
	_v0.db = db
	_v0.cache = cache
	_v0.total(group)
}

func (s service) total(group *gin.RouterGroup) {
	group.GET("/show", cache.CachePage(s.cache, time.Minute, func(ctx *gin.Context) {
		page := model.Page(ctx.Request, new([]model.Media))

		find, err := page.Find(s.db.Table("media"))
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	}))
}
