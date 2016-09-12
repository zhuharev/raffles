package models

import (
	"time"
)

type Source struct {
	Id    int64
	Title string

	VkId   int64 `xorm:"unique"`
	LastId int

	IsPremium  bool
	PremiumEnd time.Time

	ScanInterval time.Duration

	CreatedBy int64     // кем создана
	Created   time.Time `xorm:"created"`
}

func SourcesCreate(source *Source) error {
	_, e := x.Insert(source)
	return e
}

func SourcesList() ([]*Source, error) {
	var (
		res []*Source
	)
	e := x.Find(&res)
	return res, e
}

func SourcesGetById(id int64) (*Source, error) {
	s := new(Source)
	_, e := x.Id(id).Get(s)
	return s, e
}

func SourcesHas(id int64) (bool, error) {
	s := new(Source)
	return x.Where("vk_id = ?", id).Get(s)
}
