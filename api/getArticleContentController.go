package api

import (
	"blogone/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArticleContent(ctx *gin.Context) {
	articleID := ctx.Param("article_id")
	intArticleID, err := strconv.Atoi(articleID)
	if err != nil {
		ctx.JSON(403, gin.H{
			"err":    err,
			"reason": "请输入正确的文章ID",
		})
	}
	result, err := service.GetarticleContentSrv(intArticleID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"err":    err,
			"reason": "server error.",
		})
	}

	ctx.String(200, "html", string(result))
}
