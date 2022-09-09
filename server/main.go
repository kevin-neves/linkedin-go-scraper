package main

import (
	"fmt"
	"log"

	"github.com/kevin-neves/linkedin-go-scraper/server/config"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/linkedin-go-scraper/server/jobs"
)

func getJobs(c *gin.Context) {
	jobsResponse := jobs.GetNewJobs()
	c.JSON(200, jobsResponse)
}

func getConfiguration(c *gin.Context) {
	c.JSON(200, config.GetConfiguration)
}

func postConfiguration(c *gin.Context) {
	var newConfig config.Configuration
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		log.Println(err)
		c.JSON(400, fmt.Sprintf("{error: %s}", err))
		return
	}
	config.SetConfiguration(newConfig)
	c.JSON(201, newConfig)
}

func main() {
	config.InitConfig()
	s := gin.Default()

	s.GET("/config", getConfiguration)
	s.POST("/config", postConfiguration)
	s.GET("/", getJobs)

	go jobs.UpdateJobs()

	s.Run(":8080")
}
