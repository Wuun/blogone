package service

import (
	"blogone/conf"
)
func CreateDatabase() {
	conf.MYSQL_CONNECT.AutoMigrate()
}
