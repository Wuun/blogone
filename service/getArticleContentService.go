package service

import (
	"bblog/conf"
	"bblog/serializer"
	"log"
)

type ArticleContentSrv struct {
	Title    string
	Content  string
	Comments []serializer.Comment
}

func (artContentResp *ArticleContentSrv) GetContent(articleId string) (err error) {
	var (
		content  serializer.ArticleContent
		article  serializer.ArticleModel
		comments []serializer.Comment
	)

	if err = conf.MYSQL_CONNECT.Where("article_model_id = ?", articleId).Find(&content).Error; err != nil {
		log.Print(err)
		return err
	}
	if err = conf.MYSQL_CONNECT.Where("uuid = ?", articleId).Find(&article).Error; err != nil {
		log.Print(err)
		return
	}
	if err = conf.MYSQL_CONNECT.Where("article_id = ?", articleId).Find(&comments).Error; err != nil {
		log.Print(err)
		return
	}
	artContentResp.Content = content.Content
	artContentResp.Title = article.Title
	artContentResp.Comments = comments
	return
}
