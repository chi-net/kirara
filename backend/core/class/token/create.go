package token

import (
	"github.com/chi-net/kirara/core/db/sqlite"
	"github.com/chi-net/kirara/core/utils"
	"strconv"
)

func Create(uid int) string {
	db := sqlite.GetDatabaseInstance()

	token := utils.GenerateRandomString(32)

	resp, err := db.Exec("INSERT INTO `caches`(`cid`,`key`,`value`) VALUES (NULL, ?, ?);", "appsessiontoken:"+token, strconv.Itoa(uid))

	if err != nil {
		return ""
	}

	if res, err := resp.RowsAffected(); err != nil || res != 1 {
		return ""
	}

	return token
}
