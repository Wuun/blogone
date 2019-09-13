package api

import (
	"bblog/serializer"
	"bblog/service"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ListArticle list all article exist in the server.
func ListArticle(ctx *gin.Context) {
	hpSrv := service.HomePageServ{}
	limit, err := strconv.Atoi(os.Getenv("WEB_LIMIT"))
	if err != nil {
		hpSrv.SetLimit(4)
	} else {
		hpSrv.SetLimit(limit)
	}
	rawOffset := ctx.Query("offset")
	offset, err := strconv.Atoi(rawOffset)
	if err != nil {
		offset = 0
	}
	hpSrv.SetOffset(offset)
	hpSrv.List()
	ctx.HTML(200, "article.html", hpSrv)
}

//GetArticleContent get article content of specific article.
func GetArticleContent(ctx *gin.Context) {
	articleID := ctx.Param("article_id")
	content := &service.ArticleContentSrv{}
	err := content.GetContent(articleID)
	if err != nil {
		ctx.JSON(200, serializer.BuildResponse(40002, "no such article", nil, err))
		return
	}

	ctx.HTML(200, "article_content.html", content)
}

//UploadArticle , this api is use to upload article
func UploadArticle(ctx *gin.Context) {
	uploadArticleSrv := service.UploadArticleSrv{}
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(200, serializer.Response{
			Status: 40001,
			Msg:    "error when get your data.",
			Error:  err,
		})
		return
	}
	err = json.Unmarshal(body, &uploadArticleSrv)
	if err != nil {
		ctx.JSON(200, serializer.Response{
			Status: 40001,
			Msg:    "error when unmarshal your data.",
			Error:  err,
		})
		return
	}
	resp := uploadArticleSrv.Upload()
	ctx.JSON(200, resp)
}

//WriteArticle this api is use to direct to write article page.
func WriteArticle(c *gin.Context) {
	c.HTML(200, "write_article.html", nil)
}

//CommentToArticle set comment to article
func CommentToArticle(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(200, &serializer.Response{
			Status: 4001,
			Msg:    "comment failed.",
		})
		return
	}
	commentReq := serializer.CommentRequest{}
	err = json.Unmarshal(body, &commentReq)
	if err != nil {
		ctx.JSON(200, &serializer.Response{
			Status: 4001,
			Msg:    "comment failed.",
		})
		return
	}
	if commentReq.Comment == "" || commentReq.Nickname == "" {
		ctx.JSON(200, &serializer.Response{
			Status: 4001,
			Msg:    "comment failed.",
		})
		return
	}
	comment := service.ArticleCommentSrv{
		ArticleID: ctx.Param("article_id"),
		Author:    commentReq.Nickname,
		Comment:   commentReq.Comment,
	}
	ctx.JSON(200, comment.Add())
}

func ModifyInformationPage(c *gin.Context) {
	var (
		modify service.ModifyArticleSrv
	)
	id := c.Param("article_id")
	if id == "" {
		c.JSON(200, serializer.Response{
			Status: 40001,
			Msg:    "can't find article.error-1",
		})
		return
	}

	err := modify.FindInformation(id)
	if err != nil {
		c.JSON(200, serializer.Response{
			Status: 40001,
			Msg:    "can't find article.error-2",
		})
		return
	}

	c.HTML(200, "modify_article.html", modify)
}

func ModifyArticle(c *gin.Context) {
	var (
		modify service.ModifyArticleSrv
	)
	id := c.Query("article_id")
	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		c.JSON(200, serializer.Response{
			Status: 40001,
			Error:  err,
		})
		return
	}

	err = json.Unmarshal(body, &modify)
	if err != nil {
		c.JSON(200, serializer.Response{
			Status: 40001,
			Error:  err,
		})
		return
	}
	err = modify.Modify(id)
	if err != nil || modify.Content == "" || modify.Title == "" || modify.Description == "" {
		c.JSON(200, serializer.Response{
			Status: 40001,
			Msg:    "modify failed.",
		})
		return
	}

	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "modify successfully.",
	})
}
