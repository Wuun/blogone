package service

import (
	"bblog/conf"
	"bblog/serializer"
	"log"
)

//ListArticleSrv is use to list article.
func ListArticleSrv() (result []serializer.ArticleModel, err error) {
	if err = conf.MYSQL_CONNECT.Order("created_at desc").Find(&result).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return
}
