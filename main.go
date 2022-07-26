package main

import (
	"github.com/maxvidenin/banking-lib/logger"
	"github.com/maxvidenin/banking/app"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
