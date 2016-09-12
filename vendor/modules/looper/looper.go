package looper

import (
	"fmt"
	"models"
	"modules/setting"
	"modules/vk"
	"time"

	"github.com/Unknwon/com"
)

type looper struct {
}

func Start() error {
	go func() {

		e := checkConfigSources()
		if e != nil {
			fmt.Println(e)
			return
		}

		ticker := time.NewTicker(65 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("[ticker] Get new tasks\n")
				e = getNewPosts()
				if e != nil {
					fmt.Println(e)
				}
			}
		}

	}()
	return nil
}

func doTasks() error {
	return nil
}

func getNewPosts() error {
	list, e := models.SourcesList()
	if e != nil {
		return e
	}
	for _, v := range list {
		posts, e := vk.WallGet(v.VkId)
		if e != nil {
			return e
		}
		lid := 0
		for _, post := range posts {
			if v.LastId == 0 || post.Id > v.LastId {
				_, e := models.CreateFromVkPost(post, v.Id)
				if e != nil {
					fmt.Println("[task error]", e.Error())
				}
				if lid == 0 || lid < post.Id {
					lid = post.Id
				}
			}
		}
		v.LastId = lid
		e = models.Save(v)
		if e != nil {
			return e
		}
	}
	return nil
}

func checkConfigSources() error {
	var (
		strIds = setting.Config.App.DefaultGroups
	)

	for _, v := range strIds {
		id := com.StrTo(v).MustInt64()
		has, e := models.SourcesHas(id)
		if e != nil {
			return e
		}
		if !has {
			s := new(models.Source)
			s.VkId = id
			e = models.SourcesCreate(s)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
