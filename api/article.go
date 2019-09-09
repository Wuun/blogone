package api

import (
	"blogone/serializer"
	"blogone/service"
	"github.com/gin-gonic/gin"
)

//get article content of specific article.
func GetArticleContent(ctx *gin.Context) {
	articleID := ctx.Param("article_id")
	result, err := service.GetarticleContentSrv(articleID)
	if err != nil {
		ctx.JSON(200, serializer.BuildResponse(40002,"no such article",nil,err))
		return
	}

	ctx.String(200, "html", string(result))
}

//list all article sxit in the server.
func ListArticle(ctx *gin.Context) {
	result, err := service.ListArticleSrv()
	if err != nil {
		ctx.String(500, "", "服务器错误")
		return
	}
	ctx.HTML(200, "article.html", result)
}

//write article into mysql.
func WriteArticle(ctx *gin.Context){
	uploadArticleSrv := service.UploadArticleSrv{}
	err := ctx.ShouldBind(uploadArticleSrv)
	if err != nil {
		ctx.JSON(200,gin.H{
			"status":"4002",
			"msg":"can't not write article to server.",
		})
		return
	}
	resp := uploadArticleSrv.Upload()
	ctx.JSON(200,resp)
}

//set comment to article
func CommentToArticle(ctx *gin.Context){
	comment := service.ArticleCommentSrv{}
	if err := ctx.ShouldBind(comment);err != nil {
		ctx.JSON(200,&serializer.Response{
			Status:40001,
			Msg:"error,when try to get what you upload.",
			Error:err,
		})
		return
	}
	ctx.JSON(200,comment.Add())
}