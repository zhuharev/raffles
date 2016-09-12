package tasks

import (
	"github.com/gin-gonic/gin"
	"models"
)

func Feed(c *gin.Context) {
	tasks, e := models.TasksList()
	if e != nil {
		c.JSON(200, gin.H{"error": e.Error()})
		return
	}
	c.JSON(200, gin.H{
		"response": tasks,
	})
}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": nil,
	})
}
