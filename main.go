package main

import (

	//"models"
	"GinHTMLRender"
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

	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"

	r := gin.Default()
	r.HTMLRender = htmlRender.Create()
	r.Static("/public", "./public")
	//r.Use(...)

	r.Any("/edittask/:id", tasks.Edit)

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
			sourcesGr.GET("/list", sources.List)
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
	r.GET("/temp", func(c *gin.Context) {
		c.HTML(200, "test", nil)
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
