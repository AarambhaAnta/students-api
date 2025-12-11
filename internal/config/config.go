package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// find the config path in env
	configPath = os.Getenv("CONFIG_PATH")

	// find the config path in command arguments
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set.")
		}
	}

	// check if file exists at path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// good to go, read config
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	// failed to read config
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	// return config object pointer
	return &cfg
}
