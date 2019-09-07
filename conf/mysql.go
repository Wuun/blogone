package conf

import (
	"blogone/serializer"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MYSQL_CONNECT *gorm.DB

func initMySQL() {
	var (
		err error
	)
	sqlAdr := G_CONF.MSUser + ":" + G_CONF.MSPWD + `@tcp(`+ G_CONF.MSHost+":"+ G_CONF.MSPort + `)/` + G_CONF.MSDBName + "?charset=utf8"
	MYSQL_CONNECT, err = gorm.Open("mysql", sqlAdr)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect MySQL successfully.")
	AutoCreateDatabaseSrv()
	MYSQL_CONNECT.DB().SetMaxIdleConns(10)
	MYSQL_CONNECT.DB().SetMaxOpenConns(100)
}

func AutoCreateDatabaseSrv() {
	if err := MYSQL_CONNECT.AutoMigrate(&serializer.ArticleModel{}).Error; err != nil {
		panic("error,when try to create table")
	}
	if err := MYSQL_CONNECT.AutoMigrate(&serializer.ArticleContent{}).Error; err != nil {
		panic(err)
	}
}

