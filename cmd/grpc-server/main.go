package main

import (
	"context"
	"flag"
	"log"

	"github.com/Mobo140/auth/internal/app"
	"github.com/Mobo140/auth/internal/metric"
)

var (
	configPath string
	logLevel   string
)

func setupFlags() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
	flag.StringVar(&logLevel, "l", "info", "log level")
	flag.Parse()
}

func main() {
	setupFlags()
	ctx := context.Background()

	err := metric.Init(ctx)
	if err != nil {
		log.Fatalf("failed to init metric: %v", err)
	}

	a, err := app.NewApp(ctx, configPath, logLevel)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
