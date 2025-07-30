package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbHost     string `yaml:"dbHost" env-required:"true"`
	DbPort     string `yaml:"dbPort" env:"ENV" env-required:"true"`
	DbUser     string `yaml:"dbUser" env:"ENV" env-required:"true"`
	DbPassword string `yaml:"dbPassword" env:"ENV" env-required:"true"`
	DbName     string `yaml:"dbName" env:"ENV" env-required:"true"`
	HttpPort   string `yaml:"httpPort" env:"ENV" env-required:"true"`
}

func Init() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config Path not found")
		}

	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Cannot find the config file: %s", err.Error())
	}

	return &cfg

}