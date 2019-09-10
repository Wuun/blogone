package router

import (
	"bblog/api"
	"bblog/conf"
	"bblog/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	group := router.Group("/api/v1")
	group.Use(middleware.Session(conf.G_CONF.Secret))
	{
		group.GET("/login",api.Login)
		group.POST("/authenticate",api.Authenticate)
		group.GET("/ping", api.Ping)
		group.GET("/article_list", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
		group.POST("/comment/:article_id", api.CommentToArticle)

		authen := group.Group("write_article/")
		authen.Use(middleware.UploadAuth())
		{
			authen.GET("write", api.WriteArticle)
			authen.POST("upload_article", api.UploadArticle)
		}
	}

	return router
}
