package main

import "blogone/router"

func main() {
	router := router.NewRouter()
	router.Run("127.0.0.1:8948")
}
