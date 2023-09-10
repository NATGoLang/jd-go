package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`
}

func InitConfig() (Config, error) {
	f, err := os.Open("configs/private_config.yaml")
	if err != nil {
		// return empty Config here is weird 
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
