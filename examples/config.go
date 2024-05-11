package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andiksetyawan/config"
)

// get env from os
func main() {
	os.Setenv("USER_SVC_DB_NAME", "user")
	os.Setenv("USER_SVC_DB_HOST", "localhost")

	cfg := struct {
		DbName string `env:"DB_NAME"`
		DbHost string `env:"DB_HOST"`
	}{}
	if err := config.New(config.WithPrefix("USER_SVC_")).Load(&cfg); err != nil {
		log.Fatal(fmt.Errorf("failed to load env: %w", err))
	}

	//{DbName:user DbHost:localhost}
	fmt.Printf("%+v", cfg)
}
