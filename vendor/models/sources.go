package models

import (
	"time"
)

type Source struct {
	Id    int64
	Title string
	VkId  int64 `xorm:"unique"`

	IsPremium  bool
	PremiumEnd time.Time

	ScanInterval time.Duration

	CreatedBy int64 // кем создана
	Created   time.Time
}

func SourcesCreate(source *Source) error {
	_, e := x.Insert(source)
	return e
}
