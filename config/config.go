package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	DataPath string `env:"DATA-PATH" env-default:"./data" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig(path string) *Config {
	once.Do(func() {
		log.Printf("read config from: %s", path)
		instance = &Config{}
		if err := cleanenv.ReadConfig(path, instance); err != nil {
			log.Fatal(err)
		}
	})

	log.Println("data path: ", instance.DataPath)
	return instance
}
