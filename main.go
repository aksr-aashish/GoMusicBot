package main

import (
	"log"

	"github.com/joho/godotenv"

	"GoMusicBot/bot"
	"GoMusicBot/i18n"
	"GoMusicBot/tgcalls"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: " + err.Error())
	}

	err = i18n.LoadFiles()
	if err != nil {
		log.Fatal("Error loading localization files: " + err.Error())
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
