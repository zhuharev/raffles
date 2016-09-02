package vk

import (
	"fmt"

	"github.com/zhuharev/vkutil"
)

func GroupIdByScreenName(screenName string) (int64, error) {
	ot, id, e := util.UtilsResolveScreenName(screenName)
	if e != nil {
		return 0, e
	}
	if vkutil.ObjectType(ot) != vkutil.OBJECT_GROUP {
		return 0, fmt.Errorf("Object not a %s, got %s", vkutil.OBJECT_GROUP, ot)
	}
	return int64(id), nil
}
