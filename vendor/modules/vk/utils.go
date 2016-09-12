package vk

import (
	//"fmt"
	//"net/url"
	"modules/setting"
	"regexp"
)

var (
	re = regexp.MustCompile(`.*vk.com/([a-z0-9\.]{2,})#?.*`)

	//ErrorScreenNameNotFound = fmt.Errorf("%s", "Screen name not found")
)

func ParseLink(link string) (string, error) {
	arr := re.FindStringSubmatch(link)
	if len(arr) != 2 {
		return link, nil
	}
	return arr[1], nil
}

func GetValidToken() (string, error) {
	return setting.Config.Vk.AccessToken, nil
}
