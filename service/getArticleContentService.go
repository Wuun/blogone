package service

import (
	"bblog/conf"
	"bblog/serializer"
	"log"
)

//ArticleContentSrv is service of article content.
type ArticleContentSrv struct {
	Title        string
	Content      string
	Comments     []serializer.Comment
	CommentCount int
}

//GetContent is use to get article cnotent.
func (artContentResp *ArticleContentSrv) GetContent(articleID string) (err error) {
	var (
		content  serializer.ArticleContent
		article  serializer.ArticleModel
		comments []serializer.Comment
	)

	if err = conf.MYSQL_CONNECT.Where("article_model_id = ?", articleID).Find(&content).Error; err != nil {
		log.Print(err)
		return err
	}
	if err = conf.MYSQL_CONNECT.Where("uuid = ?", articleID).Find(&article).Error; err != nil {
		log.Print(err)
		return
	}
	if err = conf.MYSQL_CONNECT.Where("article_id = ?", articleID).Find(&comments).Error; err != nil {
		log.Print(err)
		return
	}
	artContentResp.Content = content.Content
	artContentResp.Title = article.Title
	artContentResp.Comments = comments
	artContentResp.CommentCount = len(comments)
	return
}
