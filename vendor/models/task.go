package models

import (
	//"encoding/json"
	"time"
)

type Task struct {
	Id          int64     `json:"id"`
	SourceId    int64     `json:"source_id"`
	OwnerId     int64     `json:"owner_id"`
	PostId      int64     `json:"post_id"`
	Description string    `json:"description"`
	TimeStart   time.Time `json:"time_start"`
	TimeRemove  time.Time `json:"time_remove"`
	Time        time.Time `json:"time"`
	Actions     []Action  `xorm:"-" json:"actions"`
}

/*func (t Task) MarshalJSON() ([]byte, error) {
return json.Marshal(struct{
	Id int64
	}{})
}*/
