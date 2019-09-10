package api

import (
	"bblog/serializer"
	"bblog/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

//list all article exist in the server.
func ListArticle(ctx *gin.Context) {
	result, err := service.ListArticleSrv()
	if err != nil {
		ctx.String(200, "string", "服务器错误")
		return
	}
	ctx.HTML(200, "article.html", result)
}

//get article content of specific article.
func GetArticleContent(ctx *gin.Context) {
	articleID := ctx.Param("article_id")
	fmt.Println(articleID)
	content := &service.ArticleContentSrv{}
	err := content.GetContent(articleID)
	if err != nil {
		ctx.JSON(200, serializer.BuildResponse(40002, "no such article", nil, err))
		return
	}

	ctx.HTML(200, "article_content.html", content)
}

//write article into mysql.
func WriteArticle(ctx *gin.Context) {
	uploadArticleSrv := service.UploadArticleSrv{}
	err := ctx.ShouldBind(uploadArticleSrv)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status": "4002",
			"msg":    "can't not write article to server.",
		})
		return
	}
	resp := uploadArticleSrv.Upload()
	ctx.JSON(200, resp)
}

//set comment to article
func CommentToArticle(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(200, &serializer.Response{
			Status: 4001,
			Msg:    "comment failed.",
		})
	}
	commentReq := serializer.CommentRequest{}
	err = json.Unmarshal(body, &commentReq)
	if err != nil {
		ctx.JSON(200, &serializer.Response{
			Status: 4001,
			Msg:    "comment failed.",
		})
	}
	comment := service.ArticleCommentSrv{
		ArticleID: ctx.Param("article_id"),
		Author:    commentReq.Nickname,
		Comment:   commentReq.Comment,
	}
	ctx.JSON(200, comment.Add())
}
