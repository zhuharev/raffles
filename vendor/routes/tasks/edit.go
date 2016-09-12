package tasks

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"fmt"
	"models"
)

func Edit(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	task, e := models.TasksGet(id)
	if e != nil {
		fmt.Println(e)
	}
	c.HTML(200, "edit", task)
}
