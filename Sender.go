package main

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
)

func sendMessage(msg string) bool {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if _, err := bot.PushMessage(os.Getenv("SYAHRUL"), linebot.NewTextMessage(msg)).Do(); err != nil {
		panic(err)
		return false
	}
	return true
}