package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"io/ioutil"
	"log"
	"os"
)

 type Secret struct {
	 CHANNEL_TOKEN string `json:"CHANNEL_TOKEN"`
	 CHANNEL_SECRET string `json:"CHANNEL_SECRET"`
	 BLESSED string `json:"BLESSED"`
	 SYAHRUL string `json:"SYAHRUL"`
}

func readSecret() Secret {
	var secret Secret
    jsonFile, err := os.Open("secret.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal([]byte(byteValue), &secret)
    return secret
}

func sendMessage(msg string) bool {
	var secret = readSecret()
	bot, err := linebot.New(
		secret.CHANNEL_SECRET,
		secret.CHANNEL_TOKEN,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if _, err := bot.PushMessage(secret.SYAHRUL, linebot.NewTextMessage(msg)).Do(); err != nil {
		panic(err)
		return false
	}
	return true
}