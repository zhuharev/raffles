package models

import (
	//"encoding/json"
	"fmt"
	"github.com/zhuharev/vkutil"
	"time"
)

type Task struct {
	Id       int64 `json:"id"`
	SourceId int64 `json:"source_id"`
	OwnerId  int   `json:"owner_id"`
	PostId   int   `json:"post_id"`

	Description   string    `json:"description"`
	DoTime        time.Time `json:"time_start"`
	UnDoTime      time.Time `json:"time_remove"`
	EventTime     time.Time `json:"time"`
	PublishedTime time.Time `json:"published_time"`
	Actions       []*Action `json:"actions"`
}

func TasksHas(ownerId int, postId int) (bool, error) {
	t := new(Task)
	return x.Where("owner_id = ?", ownerId).And("post_id == ?", postId).Get(t)
}

func CreateFromVkPost(post *vkutil.Post, sourceIds ...int64) (*Task, error) {
	task := new(Task)
	task.Description = post.Text
	task.OwnerId = post.OwnerId
	task.PostId = post.Id
	task.PublishedTime = time.Time(post.Date)
	if sourceIds != nil {
		task.SourceId = sourceIds[0]
	}

	if has, e := TasksHas(post.OwnerId, post.Id); !has && e == nil {
		e := Save(task)
		return task, e
	} else {
		if e != nil {
			return nil, e
		} else {
			return nil, fmt.Errorf("Task already created")
		}
	}
	return nil, fmt.Errorf("Task already created")
}

func TasksList() ([]*Task, error) {
	var (
		res []*Task
	)
	e := x.Find(&res)
	return res, e
}

func TasksGet(id int64) (*Task, error) {
	t := new(Task)
	_, e := x.Id(id).Get(t)
	return t, e
}

func TasksGetForDo() ([]*Task, error) {
	var (
		res []*Task
	)
	e := x.Where("do_time < ", time.Now()).Find(&res)
	return res, e
}

/*func (t Task) MarshalJSON() ([]byte, error) {
return json.Marshal(struct{
	Id int64
	}{})
}*/
