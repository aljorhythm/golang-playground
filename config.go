package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DigitalOceanStoreConfigClient struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
}

type DigitalOceanStoreConfig struct {
	Client *DigitalOceanStoreConfigClient `yaml:"client"`
	Space  string                         `yaml:"space"`
}

type Config struct {
	DigitalOceanStore *DigitalOceanStoreConfig `yaml:"digital_ocean_store"`
}

func readConfig() Config {
	dat, err := os.ReadFile(".config.yml")

	if err != nil {
		log.Fatal("unable to read config.yml")
	}

	config := Config{}
	err = yaml.Unmarshal(dat, &config)

	if err != nil {
		log.Fatal("unable to parse .config.yml")
	}
	return config
}
