package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func hindex(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/main")
}

func hmain(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":       "Информация о погоде",
		"placeholder": "Введите название города...",
		"label":       "Выберите ближайшее локация:",
	})
}

func hweather(c *gin.Context) {

	var (
		requestType   = c.Param("reqType")
		cityName      = c.Query("cityName")
		apiUrl        = fmt.Sprintf("http://api.weatherapi.com/v1/%v.json?key=6ca11730e8a04feb9bf83407222606&lang=ru&q=%v&aqi=no", requestType, cityName)
		timeOutClinet = http.Client{
			Timeout: time.Second * 2,
		}
		HTTPReq, _     = http.NewRequest(http.MethodGet, apiUrl, nil)
		HTTPRes, _     = timeOutClinet.Do(HTTPReq)
		HTTPResBody, _ = ioutil.ReadAll(HTTPRes.Body)
	)

	if requestType == "search" {

		searchStruct := weatherSearch{}

		jsonErr := json.Unmarshal(HTTPResBody, &searchStruct)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		c.JSON(200, searchStruct)

	} else if requestType == "current" {

		currentStruct := weatherCurrent{}

		jsonErr := json.Unmarshal(HTTPResBody, &currentStruct)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		c.JSON(200, currentStruct)

	} else {
		c.JSON(404, gin.H{"code": "404", "message": "Undefined query"})
	}

}
