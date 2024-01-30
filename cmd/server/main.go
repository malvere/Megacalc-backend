package main

import (
	"flag"
	"log"
	"tg-backend/internal/config"
	"tg-backend/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/local.yaml", "path to config file")
}
func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {

		if err != nil {
			log.Fatal(err)
		}
	}
	if err := server.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
