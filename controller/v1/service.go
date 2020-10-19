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
	_v0.show(group)
	_v0.query(group)
}

// Show godoc
// @Summary Show data information
// @Description get all data info from server
// @Param page query string false "give your selected page"
// @Param per_page query string false "give your want show lists number on per page"
// @Produce  json
// @Success 200 {object} model.Paginator{data=[]model.Media}
// @Router /show [get]
func (s service) show(group *gin.RouterGroup) {
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

// Show godoc
// @Summary Query data information
// @Description get all data info from server
// @Accept x-www-form-urlencoded
// @Param video_no formData string false "search from video number"
// @Param intro formData string false "search from intro"
// @Param hash formData string false "search with hash code"
// @Param page query string false "give your selected page"
// @Param per_page query string false "give your want show lists number on per page"
// @Produce  json
// @Success 200 {object} model.Paginator{data=[]model.Media}
// @Router /query [post]
func (s service) query(group *gin.RouterGroup) {
	group.POST("/query", func(ctx *gin.Context) {
		page := model.Page(ctx.Request, new([]model.Media))
		m := s.db.Model(model.Media{})

		//todo: add more query arguments
		if ctx.PostForm("video_no") != "" {
			infos := s.db.Model(model.Info{}).Where("video_no = (?)", ctx.PostForm("video_no")).Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("intro") != "" {
			infos := s.db.Model(model.Info{}).Where("intro like (?)", "%"+ctx.PostForm("intro")+"%").Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("hash") != "" {
			files := s.db.Model(model.File{}).Where("root_hash = (?)", ctx.PostForm("hash")).
				Or("thumb_hash = (?)", ctx.PostForm("hash")).
				Or("poster_hash = (?)", ctx.PostForm("hash")).
				Or("source_hash = (?)", ctx.PostForm("hash")).
				Or("m3u8_hash = (?)", ctx.PostForm("hash")).
				Select("id")
			m = m.Where("media.file_id in (?)", files)
		}

		find, err := page.Find(m)
		if err != nil {
			log.Errorw("find data error", "error", err)
			FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	})
}
