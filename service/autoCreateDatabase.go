package service

import (
	"blogone/conf"
	"blogone/serializer"
)

func AutoCreateDatabaseSrv() {
	var (
		articleTable        = &serializer.ArticleModel{}
		articleContentTable = &serializer.ArticleContent{}
	)
	if err := conf.MYSQL_CONNECT.Exec("create database" + conf.G_CONF.MSDBName).Error; err != nil {
		panic(err)
	}
	if err := conf.MYSQL_CONNECT.AutoMigrate(articleTable); err != nil {
		panic(err)
	}
	if err := conf.MYSQL_CONNECT.AutoMigrate(articleContentTable); err != nil {
		panic(err)
	}
}
