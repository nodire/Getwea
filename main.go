package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()
	app.Static("/files", "./static")
	app.LoadHTMLGlob("templates/**/*")
	app.GET("/", hindex)
	app.GET("/main", hmain)
	app.GET("/weather/:reqType", hweather)
	err := app.Run(":7070")
	if err != nil {
		log.Fatalln("Ошибка запуска сервера")
	}

}
