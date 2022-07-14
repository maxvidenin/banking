package main

import (
	"github.com/maxvidenin/banking/app"
	"github.com/maxvidenin/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
