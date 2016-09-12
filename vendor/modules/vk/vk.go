package vk

import (
	//	"github.com/zhuharev/vk"
	"github.com/zhuharev/vkutil"
	"models"

	"sync"
)

var (
	util *vkutil.Api
	mux  sync.Mutex
)

func NewContext() error {
	util = vkutil.New()
	return nil
}

func Lock() {
	mux.Lock()
}

func Unlock() {
	mux.Unlock()
}

func Do(act *models.Action, at string) error {
	Lock()
	defer Unlock()
	util.VkApi.AccessToken = at
	e := act.Do(util)
	return e
}

func UnDo(act *models.Action, at string) error {
	Lock()
	defer Unlock()
	util.VkApi.AccessToken = at
	e := act.UnDo(util)
	return e
}

func WallGet(id int64) ([]*vkutil.Post, error) {
	Lock()
	defer Unlock()
	tok, _ := GetValidToken()
	util.VkApi.AccessToken = tok
	return util.WallGet(-int(id))
}
