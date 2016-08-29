package models

import (
	"github.com/zhuharev/vkutil"
)

type ActionName string

const (
	AN_Repost ActionName = "repost"
	AN_Like              = "like"
	AN_Join              = "join"
	AN_Follow            = "follow"

	AN_Delete   = "delete"
	AN_Unlike   = "unlike"
	AN_Leave    = "leave"
	AN_Unfollow = "unfollow"
)

type Action struct {
	Name          ActionName        `json:"name"`
	ObjectType    vkutil.ObjectType `json:"object_type"`
	ObjectId      int64             `json:"object_id"`
	ObjectOwnerId int64             `json:"object_owner_id"`
}
