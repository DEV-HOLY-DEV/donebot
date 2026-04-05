package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/c3k4ah/donebot/internal/notifier"
	"github.com/c3k4ah/donebot/internal/runner"
	"github.com/joho/godotenv"
)

func main() {
	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: donebot [options] <command> [args...]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Flags definition
	discordFlag := flag.Bool("dis", false, "Send a Discord notification once command finishes")
	
	// Parse flags
	flag.Parse()

	// Remaining arguments after flags
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Load .env (optional, doesn't fail if file is missing)
	_ = godotenv.Load()

	// Execution
	commandName := args[0]
	commandArgs := args[1:]
	
	err := runner.RunCommand(commandName, commandArgs...)
	if err != nil {
		log.Printf("Execution error: %v", err)
	}

	// Notification
	if *discordFlag {
		webhookURL := os.Getenv("WEBHOOK_URL")
		if webhookURL == "" {
			log.Fatal("Error: WEBHOOK_URL environment variable is not set")
		}

		errNotify := notifier.SendDiscordNotification(webhookURL, commandName, err == nil)
		if errNotify != nil {
			log.Printf("Notification error: %v", errNotify)
		} else {
			fmt.Println("Notification sent to Discord!")
		}
	}

	if err != nil {
		os.Exit(1)
	}
}
