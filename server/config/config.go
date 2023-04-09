package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	Config struct {
		Environment   string `yaml:"environment"`
		ListenAddress string `yaml:"listen_address"`
	}
)

func Load(path string) (*Config, error) {
	p, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(p, &cfg)

	return &cfg, err
}
