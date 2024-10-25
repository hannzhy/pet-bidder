package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int64  `env:"PORT" envDefault:"8080"`
}

func ParseConfig() (*Config, error) {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		log.Println("Error on parse config: ", err)
		return nil, err
	}

	return conf, nil
}
