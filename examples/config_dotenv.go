package main

import (
	"fmt"
	"log"

	"github.com/andiksetyawan/config"
)

// get env from .env file
func main() {
	cfg := struct {
		DbName string `env:"DB_NAME"`
		DbHost string `env:"DB_HOST"`
	}{}
	if err := config.New(config.WithEnvPath("./examples/.env")).Load(&cfg); err != nil {
		log.Fatal(fmt.Errorf("failed to load env: %w", err))
	}

	//{DbName:user DbHost:example.com}
	fmt.Printf("%+v", cfg)
}
