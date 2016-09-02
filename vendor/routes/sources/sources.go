package sources

import (
	"github.com/gin-gonic/gin"
	"models"
	//"fmt"
	"modules/vk"
)

func Add(c *gin.Context) {
	var (
		link, _ = c.GetPostForm("link")
	)
	screenName, e := vk.ParseLink(link)
	if e != nil {
		c.JSON(200, gin.H{"error": e.Error(), "input": link})
		return
	}
	id, e := vk.GroupIdByScreenName(screenName)
	if e != nil {
		c.JSON(200, gin.H{"error": e.Error(), "input": link})
		return
	}

	src := models.Source{VkId: id}

	e = models.SourcesCreate(&src)
	if e != nil {
		c.JSON(200, gin.H{"error": e.Error(), "input": link})
		return
	}

	c.JSON(200, gin.H{"response": gin.H{"screen_name": screenName, "id": id}})
}
