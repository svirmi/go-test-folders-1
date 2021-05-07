package main

import (
	"folders-one/app"
	"folders-one/logger"
)

func main() {
	logger.Info("Starting the application ...")
	app.Start()
}
