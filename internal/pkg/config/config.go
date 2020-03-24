package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	UseProxy     bool   `json:"use_proxy"`
	ProxyUrl     string `json:"proxy_url"`
	DownloadPath string `json:"download_path"`
	MediaDomain  string `json:"media_domain"`
}


var config *Config

func init() {
	GetConfig()
}

func GetConfig() *Config {
	if config == nil {
		config = readConfig()
	}
	return config
}

func readConfig() *Config {

	configFilePath := "./config.json"

	args := os.Args[1:]

	if len(args) == 2 && args[0] == "-c" {
		if path := strings.Trim(args[1], " "); path != "" {
			configFilePath = path
		}

	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalln(
			`
read config json file failure

Usage: -c, config file path

config json file content:

{
  "use_proxy": false,
  "proxy_url": "",
  "download_path": "",
  "media_domain": ""
}

				`,
		)
	}
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	configFile.Close()
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}