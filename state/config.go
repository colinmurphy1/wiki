package state

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration structs
type Config struct {
	Wiki   ConfigWiki `yaml:"wiki"`
	Files  configFiles
	Server configServer `yaml:"server"`
}

type ConfigWiki struct {
	SiteTitle       string `yaml:"title"`
	SiteDescription string `yaml:"description"`
	Theme           string `yaml:"theme"`
	IndexPage       string `yaml:"index"`
	DocumentRoot    string `yaml:"document_root"`
}

type configFiles struct {
	BaseDir  string // Base directory
	ThemeDir string // Theme directory
	usersDb  string // Users db path
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
