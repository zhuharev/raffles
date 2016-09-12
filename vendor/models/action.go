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
	TaskId int64

	Name          ActionName        `json:"name"`
	ObjectType    vkutil.ObjectType `json:"object_type"`
	ObjectId      int               `json:"object_id"`
	ObjectOwnerId int               `json:"object_owner_id"`

	RepostPostId int `json:"-"`
}

func (a *Action) Do(vk *vkutil.Api) error {
	switch a.Name {
	case AN_Repost:
		postId, e := vk.WallRepost(a.ObjectType, a.ObjectOwnerId, a.ObjectId)
		if e != nil {
			return e
		}
		a.RepostPostId = postId
		e = Save(a)
		if e != nil {
			return e
		}
	case AN_Like:
		e := vk.LikesAdd(a.ObjectType, a.ObjectOwnerId, a.ObjectId)
		if e != nil {
			return e
		}
	case AN_Follow:
		_, e := vk.FriendsAdd(a.ObjectId)
		if e != nil {
			return e
		}
	case AN_Join:
		e := vk.GroupsJoin(a.ObjectId)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *Action) UnDo(vk *vkutil.Api) error {
	switch a.Name {
	case AN_Repost:
		e := vk.WallDelete(0, a.RepostPostId)
		if e != nil {
			return e
		}
	case AN_Like:
		e := vk.LikesDelete(a.ObjectType, a.ObjectOwnerId, a.ObjectId)
		if e != nil {
			return e
		}
	case AN_Follow:
		e := vk.FriendsDelete(a.ObjectId)
		if e != nil {
			return e
		}
	case AN_Join:
		e := vk.GroupsLeave(a.ObjectId)
		if e != nil {
			return e
		}
	}
	return nil
}
