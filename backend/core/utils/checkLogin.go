package utils

import (
	"github.com/chi-net/kirara/core/db/sqlite"
)

func CheckLogin(username string, password string) bool {
	getUserNameSQL := "SELECT value FROM settings WHERE `key`='kirara.admin.username'"
	resp, err := sqlite.Query(getUserNameSQL)

	if err != nil {
		return false
	}

	user := ""

	if resp.Next() {
		resp.Scan(&user)
	} else {
		return false
	}
	resp.Close()

	getPasswordSQL := "SELECT value FROM settings WHERE `key`='kirara.admin.password'"
	resp, err = sqlite.Query(getPasswordSQL)

	if err != nil {
		return false
	}

	pwd := ""

	if resp.Next() {
		resp.Scan(&pwd)
	} else {
		return false
	}
	resp.Close()

	if user == username && CheckPassword(password, pwd) {
		return true
	} else {
		return false
	}
}
