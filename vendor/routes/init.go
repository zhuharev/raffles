package routes

import (
	"models"
	"modules/looper"
	"modules/setting"
	"modules/vk"
)

func GlobalInit() error {
	e := setting.LoadConfig()
	if e != nil {
		return e
	}

	e = models.NewEngine()
	if e != nil {
		return e
	}

	e = vk.NewContext()
	if e != nil {
		return e
	}

	e = looper.Start()
	if e != nil {
		return e
	}

	return nil
}
