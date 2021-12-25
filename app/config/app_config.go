package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig interface {
	GetAppConfig() *config
}

// Config denotes application config
type config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"server"`
	DB struct {
		UserName string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
	} `yaml:"db"`
}

func (c *config) GetAppConfig() *config {
	return c
}

func NewAppConfig(cfgFile string) (*config, error) {
	cfg := &config{}
	if len(cfgFile) == 0 {
		return cfg, fmt.Errorf("invalid config file %s", cfgFile)
	}
	extension := filepath.Ext(cfgFile)
	if extension == "" || extension != ".yml" {
		return cfg, fmt.Errorf("invalid file extension for file %s extension %s", cfgFile, extension)
	}
	file, err := os.Open(cfgFile)
	if err != nil {
		return cfg, fmt.Errorf("file error: %v", err)
	}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return cfg, fmt.Errorf("yaml decoder error : %v", err)
	}

	return cfg, nil
}
