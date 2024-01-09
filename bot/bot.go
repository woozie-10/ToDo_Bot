package bot

import (
	"ToDo_bot/config"
	"ToDo_bot/handlers"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	token := config.Config.GetString("bot.token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			args := update.Message.CommandArguments()
			switch cmd {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello there! Welcome to the <b>ToDO bot</b>! 🌟 I'm here to assist you with task management and boost your productivity.")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				break
			case "getTasks":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlers.GetTasks(update.Message.Chat.ID))
				msg.ParseMode = "HTML"
				bot.Send(msg)
				break

			case "addTask":
				result := handlers.AddTask(update.Message.Chat.ID, args)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				bot.Send(msg)
				break

			case "getTask":
				result := handlers.GetTaskByID(update.Message.Chat.ID, args)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				msg.ParseMode = "HTML"
				bot.Send(msg)
				break
			case "delTask":
				result := handlers.DelTaskByID(update.Message.Chat.ID, args)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				bot.Send(msg)
				break
			case "updTask":
				id := strings.Split(args, " ")[0]
				newTask := strings.Join(strings.Split(args, " ")[1:], " ")
				result := handlers.UpdTask(update.Message.Chat.ID, id, newTask)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				bot.Send(msg)
				break
			case "clearTasks":
				result := handlers.ClearTasks(update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				bot.Send(msg)
				break
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
				bot.Send(msg)
				break
			}
		}
	}
}
