package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type WebConfig struct {
	Debug     bool `yaml:"debug"`
	WebServer struct {
		Port string `yaml:"port"`
	} `yaml:"webserver"`
}

var WebConfigInstance *WebConfig

func GetWebConfigInstance() *WebConfig {
	if WebConfigInstance == nil {
		WebConfigInstance = &WebConfig{}
		WebConfigInstance.ReadConfig()
	}
	return WebConfigInstance
}

var WebConfigPath string = "config/web.yml"
var WebConfigLocalPath string = "config/web_local.yml"

func (c *WebConfig) ReadConfig() {

	configPath := WebConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = WebConfigLocalPath
		}
	}
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func (c *WebConfig) IsDebug() bool {
	return GetWebConfigInstance().Debug
}
