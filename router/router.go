package router

import (
	"blogone/api"
	"blogone/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter()*gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	group := router.Group("/api/v1")
	{
		group.GET("/ping",api.Ping)
		group.GET("/article_list", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
		group.POST("/comment/:article_id",api.CommentToArticle)
		uploadPage := group.Group("/article_upload")
		uploadPage.Use(middleware.UploadAuth())
		uploadPage.POST("/:password",api.WriteArticle)
	}

	return router
}
