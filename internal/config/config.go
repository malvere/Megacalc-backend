package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	TokenTTL time.Duration  `yaml:"token_ttl" env-required:"true"`
	Database DatabaseConfig `yaml:"database" env-required:"true"`
	Server   ServerConfig   `yaml:"server"`
	Telegram Telegram
	// JWT      jwkeys.Config  `yaml:"jwt"`
}

type ServerConfig struct {
	Port    int
	Timeout time.Duration
}

type Telegram struct {
	Token  string
	ChatId string
}
type DatabaseConfig struct {
	Driver string
	URL    string
}

func LoadConfig(cfgPath string) (*Config, error) {
	path := fetchConfigPath(cfgPath)
	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}
	cfg.Telegram.Token = os.Getenv("BOT_TOKEN")
	cfg.Telegram.ChatId = os.Getenv("CHAT_ID")
	dbURL := os.Getenv("DB_URL")
	if dbURL != "" {
		cfg.Database.URL = dbURL
	}
	return &cfg, nil

}

func fetchConfigPath(cfgPath string) string {
	var res string

	// --config="path/to/cfg.yaml"
	flag.StringVar(&res, "config", "", "path")
	flag.Parse()

	if cfgPath == "" {
		res = os.Getenv("CONFIG_PATH")
	} else {
		res = cfgPath
	}
	return res
}
