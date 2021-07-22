package main

import (
	"GoMusicBot/bot"
	"GoMusicBot/tgcalls"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: " + err.Error())
	}

	err = tgcalls.Start()
	if err != nil {
		log.Fatal("Error starting GoTGCalls: " + err.Error())
	}

	err = bot.Start()
	if err != nil {
		log.Fatal("Error starting bot: " + err.Error())
	}
}
