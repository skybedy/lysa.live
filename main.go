package main

import (
	"lysa.live/routes"
	"lysa.live/utils"
)

func main() {
	router := routes.NewRouter()
	utils.LoadTemplates("templates/*.html")
	//utils.LoadTemplates("/var/www/timechip.cz/go-www.timechip.cz/templates/*.html")
	utils.HttpServer(router)
}
