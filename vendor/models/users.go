package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id             int64 `json:"id"`
	VkId           int64 `json:"vk_id"`
	UserName       string
	HashedPassword []byte
	VkAccessToken  string
	ApiToken       string // ??????????????????????

	VkLogin    string `json:",omitempty"`
	VkPassword string `json:",omitempty"`

	Role Role // Админ, Модератор, Никто

	Created time.Time
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id   int64
		VkId int64
	}{
		u.Id,
		u.VkId,
	})
}
