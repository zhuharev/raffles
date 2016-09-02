package vk

import (
	//	"github.com/zhuharev/vk"
	"github.com/zhuharev/vkutil"
)

var (
	util *vkutil.Api
)

func NewContext() error {
	util = vkutil.New()
	return nil
}
