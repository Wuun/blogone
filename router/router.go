package router

import (
	"blogone/api"
	"github.com/gin-gonic/gin"
)

func NewRouter()*gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	router.GET("/article_list", api.ListArticle)
	router.GET("/article_content/:article_id", api.GetArticleContent)
	return router
}
