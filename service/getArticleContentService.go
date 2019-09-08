package service

import (
	"blogone/conf"
	"blogone/serializer"
	"log"
)

func GetarticleContentSrv(articleId int) (result []byte, err error) {
	var (
		articleContent serializer.ArticleContent
	)
	if err = conf.MYSQL_CONNECT.Find(&articleContent, articleId).Error; err != nil {
		log.Print(err)
		return result, err
	}

	return []byte(articleContent.Content), nil
}
