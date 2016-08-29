package models

import (
	"time"
)

type Source struct {
	Id    int64
	Title string
	VkId  int64

	IsPremium  bool
	PremiumEnd time.Time

	ScanInterval time.Duration

	CreatedBy int64 // кем создана
	Created   time.Time
}
