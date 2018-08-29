package main

import (
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/koding/multiconfig"
	log "github.com/sirupsen/logrus"
	"github.rakops.com/SD/publisher-dashboard-api/cmd/api/app"
	_ "github.rakops.com/SD/publisher-dashboard-api/cmd/api/app/docs"
)

func main() {
	// configure logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

}