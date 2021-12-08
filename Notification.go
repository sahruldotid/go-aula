package main

import (
	cron "github.com/robfig/cron/v3"
	"time"
)

 func InBetween(i, min, max int) bool {
         if (i >= min) && (i <= max) {
                 return true
         } else {
                 return false
         }
 }

func StartNotification() {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	//
	defer scheduler.Stop()

	//define task here
	scheduler.AddFunc("*/1 7-17 * * 1-5", NotifyMe)
	//start
	go scheduler.Start()

}

func NotifyMe(){
	var user Users
	if err := db.First(&user).Error; err != nil{
		return
	}
	if !checkSession(user) {
		sesskey, cookie := login(user)
		user.SessKey = sesskey
		user.Cookie = cookie[0].String()
		db.Save(&user)
	}
	events := getAttendance(user).Data.Events
	for _, event := range events {
		if InBetween(int(event.Timestart), int(time.Now().Unix()), int(time.Now().Unix())+500){
				message := "Mas rul absen\n"
				message += event.Name
				message += "\n"
				message += event.URL
				sendMessage(message)
		}
	}
}