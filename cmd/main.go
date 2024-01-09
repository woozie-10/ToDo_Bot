package main

import (
	"ToDo_bot/bot"
	"ToDo_bot/config"
	"ToDo_bot/database"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
	bot.Run()
}
