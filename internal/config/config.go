package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Service struct {
	URL string `yaml:"url"`
}

type Route struct {
	Path    string `yaml:"path"`
	Service string `yaml:"service"`
}

type Config struct {
	Services map[string]Service `yaml:"services"`
	Routes   []Route             `yaml:"routes"` //slice of routes([]Route)
}
//configuration loader
func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)  //YAML into Go structure(Unmarshal)
	if err != nil {
		return nil, err
	}

	return &config, nil
}