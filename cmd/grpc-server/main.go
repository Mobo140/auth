package main

import (
	"context"
	"flag"
	"log"

	"github.com/Mobo140/auth/internal/app"
)

var configPathM string

func setupFlags() {
	flag.StringVar(&configPathM, "config-path", ".env", "path to config file")
	flag.Parse()
}

func main() {
	setupFlags()
	ctx := context.Background()

	a, err := app.NewApp(ctx, configPathM)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
