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

func main() {
	jsonData := []byte(`{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_2_restaurant.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "https://linecorp.com"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "spacing": "md",
    "action": {
      "type": "uri",
      "uri": "https://linecorp.com"
    },
    "contents": [
      {
        "type": "text",
        "text": "Brown's Burger",
        "size": "xl",
        "weight": "bold"
      },
      {
        "type": "box",
        "layout": "vertical",
        "spacing": "sm",
        "contents": [
          {
            "type": "box",
            "layout": "baseline",
            "contents": [
              {
                "type": "icon",
                "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_regular_32.png"
              },
              {
                "type": "text",
                "text": "$10.5",
                "weight": "bold",
                "margin": "sm",
                "flex": 0
              },
              {
                "type": "text",
                "text": "400kcl",
                "size": "sm",
                "align": "end",
                "color": "#aaaaaa"
              }
            ]
          },
          {
            "type": "box",
            "layout": "baseline",
            "contents": [
              {
                "type": "icon",
                "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_large_32.png"
              },
              {
                "type": "text",
                "text": "$15.5",
                "weight": "bold",
                "margin": "sm",
                "flex": 0
              },
              {
                "type": "text",
                "text": "550kcl",
                "size": "sm",
                "align": "end",
                "color": "#aaaaaa"
              }
            ]
          }
        ]
      },
      {
        "type": "text",
        "text": "Sauce, Onions, Pickles, Lettuce & Cheese",
        "wrap": true,
        "color": "#aaaaaa",
        "size": "xxs"
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "spacer",
        "size": "xxl"
      },
      {
        "type": "button",
        "style": "primary",
        "color": "#905c44",
        "action": {
          "type": "uri",
          "label": "Add to Cart",
          "uri": "https://linecorp.com"
        }
      }
    ]
  }
}`)
container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
if err != nil {
    
}
flexMessage := linebot.NewFlexMessage("FlexWithJSON", container)

        var secret = readSecret()
        bot, err := linebot.New(
                secret.CHANNEL_SECRET,
                secret.CHANNEL_TOKEN,
        )
        if err != nil {
                log.Fatal(err)
        }
        if _, err := bot.PushMessage(secret.BLESSED, flexMessage).Do(); err != nil {
                panic(err)
        }
}


