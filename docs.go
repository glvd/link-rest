package rest

import (
	"github.com/gin-gonic/gin"
	_ "github.com/glvd/link-rest/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func apiDocs(group *gin.Engine) {
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
