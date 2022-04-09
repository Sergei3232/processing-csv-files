package structTags

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const yamlFile = "configs/config.yaml"

type Config struct {
	ImageLoader string `yaml:"imageLoader"`
	FileStorage string `yaml:"fileStorage"`
}

func NenConfig() *Config {
	var c Config
	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c
}
