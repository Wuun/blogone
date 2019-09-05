package api

import (
	"blogone/service"
	"github.com/gin-gonic/gin"
)

func ListArticle(ctx *gin.Context) {
	result, err := service.ListArticleSrv()
	if err != nil {
		ctx.String(400, "", "服务器错误")
	}
	ctx.HTML(200, "article.html", result)
}
