package main

import (
	"swagger_to_test/internal/tui"
	"swagger_to_test/pkg/logger"
)

func main() {
	logger.InitLogger("logs/app.log", logger.DEBUG)
	logger := logger.GetLogger()
	app, err := tui.NewApp()
	if err != nil {
		logger.Error("%s", err.Error())
	}
	if _, err := app.Run(); err != nil {
		logger.Fatal("TUI app failed")
	}
}
