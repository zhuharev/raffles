package main

import (
	//"models"
	"routes"
	"routes/sources"
	"routes/tasks"

	"github.com/gin-gonic/gin"
)

func main() {

	e := routes.GlobalInit()
	if e != nil {
		panic(e)
	}

	r := gin.Default()
	r.Static("/public", "./public")
	//r.Use(...)

	gr := r.Group("/api/v1")
	{

		tasksGr := gr.Group("/tasks")
		{
			tasksGr.GET("/feed", tasks.Feed)
			tasksGr.GET("/myfeed", tasks.Feed)
			tasksGr.GET("/get/:id", tasks.Get)
			tasksGr.POST("/edit/:id")
		}

		sourcesGr := gr.Group("/sources")
		{
			sourcesGr.GET("/list")
			sourcesGr.POST("/add", sources.Add)
			sourcesGr.GET("/get/:id")
		}

		gr.POST("/likes/add")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func feed(c *gin.Context) {
	c.JSON(200, gin.H{"da": "lol"})
}

type watchForm struct {
	Link string `json:"link" form:"link" binding:"required"`
}

func watch(c *gin.Context) {
	var form watchForm
	if e := c.Bind(&form); e == nil {
		c.JSON(200, gin.H{"response": form.Link})
		return
	} else {
		c.JSON(500, gin.H{"error": e})
		return
	}

	c.JSON(500, gin.H{"success": false})
}
