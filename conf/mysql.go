package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MYSQL_CONNECT *gorm.DB

func init() {
	var(
		err error
	)
	conf := newConf()
	sqlAdr := conf.MSUser + ":" + conf.MSPWD + "@/" + conf.MSDBName + "?charset=utf8&parseTime=True&loc=Local"
	MYSQL_CONNECT, err = gorm.Open("mysql", sqlAdr)
	if err != nil {
		panic(err)
	}
}
