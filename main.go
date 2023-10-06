package main

import (
	"ToDo_bot/api"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Initialize the database connection
	api.InitDB()

	// Telegram Bot API token
	token := os.Getenv("TOKEN")

	// Create a new Telegram bot instance
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Enable debug mode to log API requests
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up updates polling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Handle incoming updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			cmd := update.Message.Command() // Get the command (/GetTasks)
			args := update.Message.CommandArguments()

			// Handle different bot commands
			if cmd == "GetTasks" {
				// Retrieve and send tasks information
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, api.GetInfo())
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}
			if cmd == "AddTask" {
				// Add a new task
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, api.AddInfo(args))
				bot.Send(msg)
			}
			if cmd == "ClearTasks" {
				// Clear all tasks
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, api.ClearInfo())
				bot.Send(msg)
			}
			if cmd == "DelTask" {
				// Delete a task by ID
				id, err := strconv.Atoi(args)
				var msgText string
				if err != nil {
					msgText = "Error converting string to number"
				}
				msgText = api.DelInfo(id)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				bot.Send(msg)
			}
			if cmd == "UpdTask" {
				// Update a task by ID with new content
				strId := strings.Split(args, " ")[0]
				newTask := strings.Join(strings.Split(args, " ")[1:], " ")
				id, err := strconv.Atoi(strId)
				var msgText string
				if err != nil {
					msgText = "Error converting string to number"
				}
				msgText = api.UpdInfo(id, newTask)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}
		}
	}
}
