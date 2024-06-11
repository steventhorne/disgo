package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/steventhorne/disgo/core"
)

func main() {
	godotenv.Load()

	bot := core.NewBot()
	bot.Open()
	err := bot.RegisterCommands()
	for _, e := range err {
		slog.Error("Error registering command", "error", e)
	}
}
