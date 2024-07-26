package tokens

import (
	"github.com/chi-net/kirara/core/db/sqlite"
	"github.com/chi-net/kirara/core/utils"
	"strconv"
)

func Create(uid int) string {
	token := utils.GenerateRandomString(32)

	// avoid conflict
	for Query(token) != -1 {
		token = utils.GenerateRandomString(32)
	}
	resp, err := sqlite.Exec("INSERT INTO `caches`(`cid`,`key`,`value`) VALUES (NULL, ?, ?);", "appsessiontoken:"+token, strconv.Itoa(uid))

	if err != nil {
		return ""
	}

	if res, err := resp.RowsAffected(); err != nil || res != 1 {
		return ""
	}

	return token
}
