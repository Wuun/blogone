package main

import (
	"bblog/conf"
	"bblog/router"
)

func main() {
	router := router.NewRouter()
	router.Run(conf.G_CONF.WebAddr)
}
