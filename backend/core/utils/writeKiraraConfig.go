package utils

import (
	"encoding/json"
	"os"
)

type KiraraConfig struct {
	DBPath     string `json:"DBPath"`
	ListenPort int    `json:"ListenPort"`
}

func WriteKiraraConfig(dbPath string, listenPort int) error {
	dir, _ := os.Getwd()
	path := dir + string(os.PathSeparator) + "kirara.config.json"

	jsonData, err := json.Marshal(KiraraConfig{dbPath, listenPort})
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
