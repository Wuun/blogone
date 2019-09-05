package service

import (
	"blogone/conf"
	"blogone/serializer"
	"log"
)

func ListArticleSrv() (result []serializer.ArticleModel, err error) {
	var (
		articleList []serializer.ArticleModel
	)

	if err = conf.MYSQL_CONNECT.Find(&articleList).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return
}
