package Client

import (
	"encoding/json"
	"fmt"
	"os"

	Declared "github.com/XanaOG/Cleaner/Core/Declared"
)

var (
	ConfigFile = "Assets/Config.json"
)

func GetConfig(file string) Declared.Config {
	var config Declared.Config
	ConfigFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Config file not found." + err.Error())
	}
	defer ConfigFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	JSONParser := json.NewDecoder(ConfigFile)
	JSONParser.Decode(&config)
	return config
}
