package main

import (
	"log"
	"net/http"
	"os"

	"github.com/1ef7yy/brand_scout/internal/api/routes"
	"github.com/1ef7yy/brand_scout/internal/view"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

func main() {
	logger := logger.NewLogger()

	view, err := view.NewView(logger)

	if err != nil {
		log.Fatalf("error initializing view: %s", err.Error())
	}

	mux := routes.InitHandlers(view)

	serverAddr, ok := os.LookupEnv("SERVER_ADDRESS")

	if !ok {
		serverAddr = "localhost:8080"
		logger.Infof("could not resolve SERVER_ADDRESS from env, reverting to default: %s", serverAddr)
	}

	logger.Infof("starting service at %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		logger.Fatalf("service failed: %s", err.Error())
	}
}
