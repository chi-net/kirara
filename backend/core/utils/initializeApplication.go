package utils

import (
	"errors"
	"github.com/chi-net/kirara/core/db/sqlite"
	"os"
)

func InitializeApplication(username string, password string, listenPort int, dbtype string, bottoken string) error {
	dir, _ := os.Getwd()

	if dbtype == "SQLite" {
		dbpath := dir + string(os.PathSeparator) + "kirara.app.db"
		err := sqlite.New(dbpath)
		if err != nil {
			return err
		}

		// creating tables
		_, err = sqlite.Exec("CREATE TABLE `messages` (`mid`INTEGER,`id`INTEGER,`from`INTEGER,`chat`INTEGER,`date`INTEGER,`content`TEXT,`quote`INTEGER,`type`INTEGER,PRIMARY KEY(`mid`))")
		if err != nil {
			return err
		}

		_, err = sqlite.Exec("CREATE TABLE `users` (`uid`INTEGER,`id`INTEGER,`isbot`INTEGER,`name`TEXT,`username`TEXT,PRIMARY KEY(`uid` AUTOINCREMENT))")
		if err != nil {
			return err
		}

		_, err = sqlite.Exec("CREATE TABLE `chats` (`cid`INTEGER,`id`INTEGER,`type`TEXT,`title`TEXT,`username`TEXT,PRIMARY KEY(`cid` AUTOINCREMENT))")
		if err != nil {
			return err
		}

		_, err = sqlite.Exec("CREATE TABLE `settings` (`sid`INTEGER,`key`TEXT,`value`TEXT,PRIMARY KEY(`sid` AUTOINCREMENT))")
		if err != nil {
			return err
		}

		_, err = sqlite.Exec("CREATE TABLE `caches` (`cid`INTEGER,`key`TEXT,`value`TEXT,PRIMARY KEY(`cid` AUTOINCREMENT))")
		if err != nil {
			return err
		}

		// write kirara.config.json
		WriteKiraraConfig(dbpath, listenPort)

		// add username and password into settings
		result, err := sqlite.Exec("INSERT INTO settings (key, value) VALUES (`kirara.admin.username`, `" + username + "`);")
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		if affected == 0 {
			return errors.New("username insert failed")
		}

		cryptedPassword, _ := saltPassword([]byte(password), 64)
		result, err = sqlite.Exec("INSERT INTO settings (key, value) VALUES (`kirara.admin.password`, `" + string(cryptedPassword) + "`);")
		if err != nil {
			return err
		}
		affected, _ = result.RowsAffected()
		if affected == 0 {
			return errors.New("password insert failed")
		}

		// insert telegram bot token
		result, err = sqlite.Exec("INSERT INTO settings (key, value) VALUES (`kirara.bot.token`, `" + bottoken + "`);")
		if err != nil {
			return err
		}
		affected, _ = result.RowsAffected()
		if affected == 0 {
			return errors.New("bot token insert failed")
		}
	}
	return nil
}
