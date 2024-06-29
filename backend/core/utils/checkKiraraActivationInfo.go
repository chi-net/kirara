package utils

import (
	"os"
)

func CheckKiraraActivationInfo() bool {
	dir, _ := os.Getwd()

	if CheckConfiguration(dir + string(os.PathSeparator) + "kirara.config.json") {

		conf := ReadJSONConfiguration(dir + string(os.PathSeparator) + "kirara.config.json")

		if conf.DbPath == "Failed to GET" {
			return false
		} else {
			return true
		}
	}
	return false
}
