package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MYSQL_CONNECT *gorm.DB

func init() {
	var (
		err error
	)
	sqlAdr := G_CONF.MSUser + ":" + G_CONF.MSPWD + "@/" + G_CONF.MSDBName + "?charset=utf8&parseTime=True&loc=Local"
	MYSQL_CONNECT, err = gorm.Open("mysql", sqlAdr)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect MySQL successfully.")
}
