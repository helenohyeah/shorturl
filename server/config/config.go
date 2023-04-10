package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	Config struct {
		Environment   string      `yaml:"environment"`
		ListenAddress string      `yaml:"listen_address"`
		DB            configRDBMS `yaml:"database"`
	}

	configRDBMS struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
)

func (cfg *Config) GetPostgresURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)
}

func Load(path string) (*Config, error) {
	p, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(p, &cfg)

	return &cfg, err
}
