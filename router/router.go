package router

import (
	"bblog/api"
	"bblog/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	group := router.Group("/api/v1")
	{
		group.GET("/ping", api.Ping)
		group.GET("/article_list", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
		group.POST("/comment/:article_id", api.CommentToArticle)
		uploadPage := group.Group("/write_article")
		uploadPage.Use(middleware.UploadAuth())
		uploadPage.GET("/", api.WriteArticle)
		uploadPage.POST("/upload_article", api.UploadArticle)
	}

	return router
}
