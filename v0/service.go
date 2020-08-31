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
	_v0.query(group)
}

func (s service) total(group *gin.RouterGroup) {
	group.GET("/show", cache.CachePage(s.cache, time.Minute, func(ctx *gin.Context) {
		page := model.Page(ctx.Request, new([]model.Media))

		find, err := page.Find(s.db.Model(model.Media{}))
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	}))
}

func (s service) query(group *gin.RouterGroup) {
	group.POST("/query", cache.CachePage(s.cache, time.Minute, func(ctx *gin.Context) {
		page := model.Page(ctx.Request, new([]model.Media))
		m := s.db.Model(model.Media{})
		if ctx.PostForm("video_no") != "" {
			infos := s.db.Model(model.Info{}).Where("video_no like (?)", "%"+ctx.PostForm("video_no")+"%").Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("intro") != "" {
			infos := s.db.Model(model.Info{}).Where("intro like (?)", "%"+ctx.PostForm("intro")+"%").Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("hash") != "" {
			files := s.db.Model(model.File{}).Where("thumb_hash = (?)", ctx.PostForm("hash")).
				Or("poster_hash = (?)", ctx.PostForm("hash")).
				Or("source_hash = (?)", ctx.PostForm("hash")).
				Or("m3u8_hash = (?)", ctx.PostForm("hash")).Select("id")
			m = m.Where("media.file_id in (?)", files).Or("root = (?)", ctx.PostForm("hash"))
		}

		find, err := page.Find(m)
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	}))
}
