package main

import (
	"blogone/conf"
	"blogone/router"
)

func main() {
	router := router.NewRouter()
	router.Run(conf.G_CONF.WebAddr)
}
