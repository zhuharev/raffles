package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Config struct {
		Vk struct {
			AppId       string
			AccessToken string
		}

		Db struct {
			Driver  string
			Setting string
		}
		App struct {
			DefaultGroups []string
		}
	}

	//conf *ini.File
)

func LoadConfig() error {
	iniFile, e := ini.Load("conf.ini")
	if e != nil {
		return nil
	}
	iniFile.NameMapper = ini.TitleUnderscore
	//conf = iniFile

	e = iniFile.MapTo(&Config)
	if e != nil {
		return nil
	}
	fmt.Println("[config]", Config)
	return nil
}
