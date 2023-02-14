package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DbSource       string `yaml:"DbSource"`
	EtcdAdress     string `yaml:"EtcdAdress"`
	ServiceAddress string `yaml:"ServiceAddress"`
	ServiceName    string `yaml:"ServiceName"`
}

func Parse(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
