package router

import (
	"bblog/api"
	"bblog/conf"
	"bblog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

//NewRouter is thr factory of router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	group := router.Group("/api/v1")
	group.Use(middleware.Session(conf.G_CONF.Secret))
	{
		group.StaticFS("/static", http.Dir("./view"))
		group.GET("/login", api.Login)
		group.POST("/authenticate", api.Authenticate)
		group.GET("/ping", api.Ping)
		group.GET("/home", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
		group.POST("/comment/:article_id", api.CommentToArticle)

		auth := group.Group("write_article/")
		auth.Use(middleware.Auth())
		{
			auth.GET("write", api.WriteArticle)
			auth.POST("upload_article", api.UploadArticle)
			auth.GET("modify_article_page/:article_id", api.ModifyInformationPage)
			auth.POST("modify_article/:article_id", api.ModifyArticle)
		}
	}

	return router
}
