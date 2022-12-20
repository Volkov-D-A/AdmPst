package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Volkov-D-A/AdmPst/pkg/config"
	"github.com/Volkov-D-A/AdmPst/pkg/dataserver"
	"github.com/Volkov-D-A/AdmPst/pkg/handlers"
	"github.com/Volkov-D-A/AdmPst/pkg/logs"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// get configuration
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("error getting config: %v", err)
	}

	fmt.Println(cfg)

	//Get logging instance
	logger := logs.Get(cfg.LogLevel)
	logger.Infof("Configuration loaded successfully. Logging instance initialized with log level: %v", cfg.LogLevel)

	//Create DataHandler
	dataHandler := handlers.NewDataHandler()

	//Create and initialize dataServer instance
	dataServer := new(dataserver.Server)

	go func() {
		if err := dataServer.Run(dataHandler.InitRoutes(), cfg.DataServer.Port); err != nil && err != http.ErrServerClosed {
			logger.Errorf("error running data server: %v", err)
		}
	}()

	logger.Infof("data server successfully started on port: %v", cfg.DataServer.Port)

	//Graceful shutdown data server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Info("shutting down callback server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := dataServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error while shutting down callback server")
	}

	return nil
}
