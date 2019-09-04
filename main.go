package main

import (
	"blogone/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	group := router.Group("/v1/article", nil)
	{
		group.GET("/article_list", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
	}
	err := router.Run("127.0.0.1:8948")
	if err != nil {
		panic(err)
	}
}
