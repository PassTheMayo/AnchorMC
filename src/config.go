package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	DefaultConfig *Config = &Config{
		Host: "0.0.0.0",
		Port: 25565,
	}
)

type Config struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

func (c *Config) ReadFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}

func (c *Config) WriteFile(file string) error {
	data, err := yaml.Marshal(c)

	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0777)
}
