package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/steventhorne/disgo/core"
)

func main() {
	godotenv.Load()

	bot := core.NewBot()
	bot.OpenAndRun()
	defer bot.Close()

	stchan := make(chan os.Signal, 1)
	signal.Notify(stchan, syscall.SIGTERM, os.Interrupt, syscall.SIGSEGV)
	for {
		select {
		case <-stchan:
			return
		default:
		}
		time.Sleep(1 * time.Second)
	}
}
