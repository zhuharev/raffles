package routes

import (
	"models"
	"modules/setting"
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
	return nil
}
