package state

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration structs
type Config struct {
	Wiki   ConfigWiki   `yaml:"wiki"`
	Files  configFiles  `yaml:"files"`
	Server configServer `yaml:"server"`
}

type ConfigWiki struct {
	SiteTitle       string `yaml:"title"`
	SiteDescription string `yaml:"description"`
	Theme           string `yaml:"theme"`
	IndexPage       string `yaml:"index"`
}

type configFiles struct {
	UsersDb      string `yaml:"users"`
	DocumentRoot string `yaml:"document_root"`
}

type configServer struct {
	Address string `yaml:"addr"`
	Port    uint   `yaml:"port"`
}

func (cfg *Config) ParseConfig(configFile string) error {
	// Open YAML file
	yamlData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	// Unmarshall YAML to Config struct
	if err := yaml.Unmarshal(yamlData, cfg); err != nil {
		return err
	}

	return nil
}
