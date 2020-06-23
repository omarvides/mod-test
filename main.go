package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig/pkg/envconfig"
)

func main() {
	ctx := context.Background()
	var c config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded Port: %d", c.Port)
}

type config struct {
	Port int `env:"PORT"`
}
