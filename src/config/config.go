package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Path struct {
		Images string
		Files  string
	}
	Database struct {
		TablePrefix string `yaml:"table_prefix"`
	}
	Logging struct {
		Level string
	}
}

func NewConfig() (*Config, error) {
	c := Config{}

	configPath, err := filepath.Abs(`config.yaml`)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
