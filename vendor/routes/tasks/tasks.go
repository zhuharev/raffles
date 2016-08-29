package tasks

import (
	"github.com/gin-gonic/gin"
	"models"
)

func Feed(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": []models.Task{models.ExampleTask1, models.ExampleTask2},
	})
}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": models.ExampleTask1,
	})
}
