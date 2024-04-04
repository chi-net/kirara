package utils

import (
	"os"
)

func CheckConfiguration(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
