package utils

import (
	"encoding/json"
	"os"
)

type JsonConfiguration struct {
	DbPath     string `json:"DBPath"`     // Application Database Path
	ListenPort int    `json:"ListenPort"` // Application Server Listen Port(optional)
}

func ReadJSONConfiguration(path string) JsonConfiguration {
	file, err := os.Open(path)
	if err != nil {
		return JsonConfiguration{
			DbPath:     "Failed to GET",
			ListenPort: -1,
		}
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := JsonConfiguration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return JsonConfiguration{
			DbPath:     "Failed to GET",
			ListenPort: -1,
		}
	}
	return configuration
}
