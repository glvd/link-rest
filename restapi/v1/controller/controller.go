package controller

import (
	v1 "github.com/glvd/link-rest/restapi/v1"
	"github.com/goextension/log"
	"net/http"
	"time"

	"github.com/glvd/link-rest/restapi/common/controller"
	cm "github.com/glvd/link-rest/restapi/common/model"
	"github.com/glvd/link-rest/restapi/v1/model"

	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
)

func RegisterHandle(c *controller.Controller) error {
	if err := cm.Migration(c.DB); err != nil {
		return err
	}
	group := c.Engine.Group(v1.Version)
	Show(c, group)
	Query(c, group)
	return nil
}

// Show godoc
// @Summary Show data information
// @Description get all data info from server
// @Param page query string false "give your selected page"
// @Param per_page query string false "give your want show lists number on per page"
// @Produce  json
// @Success 200 {object} model.Paginator{data=[]model.Media}
// @Router /v1/show [get]
func Show(c *controller.Controller, group *gin.RouterGroup) {
	group.GET("/show", cache.CachePage(c.Cache, time.Minute, func(ctx *gin.Context) {
		page := cm.Page(ctx.Request, new([]model.Media))
		find, err := page.Find(c.DB.Model(model.Media{}))
		if err != nil {
			log.Errorw("find data error", "error", err)
			controller.FailedJSON(ctx, "data not found")
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
// @Router /v1/query [post]
func Query(c *controller.Controller, group *gin.RouterGroup) {
	group.POST("/query", func(ctx *gin.Context) {
		page := cm.Page(ctx.Request, new([]model.Media))
		m := c.DB.Model(model.Media{})

		//todo: add more query arguments
		if ctx.PostForm("video_no") != "" {
			infos := c.DB.Model(model.Info{}).Where("video_no = (?)", ctx.PostForm("video_no")).Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("intro") != "" {
			infos := c.DB.Model(model.Info{}).Where("intro like (?)", "%"+ctx.PostForm("intro")+"%").Select("id")
			m = m.Where("media.info_id in (?)", infos)
		}

		if ctx.PostForm("hash") != "" {
			files := c.DB.Model(model.File{}).Where("root_hash = (?)", ctx.PostForm("hash")).
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
			controller.FailedJSON(ctx, "data not found")
			return
		}

		ctx.JSON(http.StatusOK, find)
	})
}
