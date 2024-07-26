package tokens

import (
	"github.com/chi-net/kirara/core/db/sqlite"
	"strconv"
)

func Query(token string) int {
	result, err := sqlite.Query("SELECT value FROM caches WHERE key = ?", "appsessiontoken:"+token)

	if err != nil {
		return -1
	}

	id := ""
	if result.Next() {
		result.Scan(&id)
		numid, _ := strconv.Atoi(id)
		return numid
	} else {
		return -1
	}

}
