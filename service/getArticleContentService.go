package service

import (
	"blogone/conf"
	"blogone/serializer"
	"log"
)

func GetArticleContentSrv(articleID int) (result serializer.ArticleContent, err error) {
	var (
		articleContent serializer.ArticleContent
	)
	if err = conf.MYSQL_CONNECT.Find(&articleContent, articleID).Error; err != nil {
		log.Print(err)
		return result, err
	}
	return
}
