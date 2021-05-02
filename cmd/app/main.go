package main

import (
	"errors"
	"log"
	"os"

	"github.com/DanielTitkov/cm-stat-engine/cmd/app/prepare"
	"github.com/DanielTitkov/cm-stat-engine/internal/app"
	"github.com/DanielTitkov/cm-stat-engine/internal/configs"
	"github.com/DanielTitkov/cm-stat-engine/internal/logger"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal("failed creating app", err)
	}

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
