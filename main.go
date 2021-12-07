package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var rootURL = "https://hebat.elearning.unair.ac.id"
var loginURL = rootURL + "/login/index.php"
var serviceURL = rootURL + "/lib/ajax/service.php?sesskey="
var sessURL = rootURL + "/badges/mybadges.php"
var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("sqlite3", "./Bot.db")
	defer db.Close()
	db.AutoMigrate(&Users{})

	r := gin.Default()
	r.GET("/user/:id/assignment", userAssignment)
	r.GET("/user/:id/attendance", userAttendance)

	go func() {
        StartNotification()
    }()
	r.Run("0.0.0.0:1212")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}

func userAssignment(c *gin.Context){
	var user Users
	var deadline []Deadline
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil{
		c.AbortWithStatus(404)
		return
	}
	if !checkSession(user) {
		sesskey, cookie := login(user)
		user.SessKey = sesskey
		user.Cookie = cookie[0].String()
		db.Save(&user)
	}

	events := getAssignments(user).Data.Events
	for _, event := range events {
		deadline = append(deadline, Deadline{event.Name, strings.Split(event.Course.Fullname, " - ")[2], parseEpoch(event.Timestart)})
	}
	c.JSON(200, deadline)
}

func userAttendance(c *gin.Context){
	var user Users
	var deadline []Deadline
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil{
		c.AbortWithStatus(404)
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
		deadline = append(deadline, Deadline{event.Name, strings.Split(event.Course.Fullname, " - ")[2], parseEpoch(event.Timestart)})
	}
	c.JSON(200, deadline)
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}
