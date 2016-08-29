package models

import (
	"time"

	"github.com/zhuharev/vkutil"
)

var (
	ExampleSource1 = Source{
		Id:    232,
		Title: "FREE 67 | БЕСПЛАТНО в Смоленске",
		VkId:  -4794508,
	}

	ExampleAction = Action{
		Name:       AN_Join,
		ObjectType: vkutil.OBJECT_GROUP,
		ObjectId:   1,
	}

	ExampleTask1 = Task{
		Id:       12323,
		SourceId: 232,
		OwnerId:  -4794508,
		PostId:   2440,
		Time:     time.Now(),
		Actions:  []Action{ExampleAction},
	}

	ExampleTask2 = Task{
		Id:       12324,
		SourceId: 232,
		OwnerId:  -4794508,
		PostId:   2438,
		Time:     time.Now(),
		Actions:  []Action{ExampleAction},
	}
)
