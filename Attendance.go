package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)

type Attendance struct {
	Data struct {
		Categoryid int `json:"categoryid"`
		Courseid   int `json:"courseid"`
		Date       struct {
			Hours     int    `json:"hours"`
			Mday      int    `json:"mday"`
			Minutes   int    `json:"minutes"`
			Mon       int    `json:"mon"`
			Month     string `json:"month"`
			Seconds   int    `json:"seconds"`
			Timestamp int    `json:"timestamp"`
			Wday      int    `json:"wday"`
			Weekday   string `json:"weekday"`
			Yday      int    `json:"yday"`
			Year      int    `json:"year"`
		} `json:"date"`
		Defaulteventcontext int `json:"defaulteventcontext"`
		Events              []struct {
			Candelete  bool        `json:"candelete"`
			Canedit    bool        `json:"canedit"`
			Categoryid interface{} `json:"categoryid"`
			Component  string      `json:"component"`
			Course     struct {
				Coursecategory  string `json:"coursecategory"`
				Courseimage     string `json:"courseimage"`
				Enddate         int    `json:"enddate"`
				Fullname        string `json:"fullname"`
				Fullnamedisplay string `json:"fullnamedisplay"`
				Hasprogress     bool   `json:"hasprogress"`
				Hidden          bool   `json:"hidden"`
				ID              int    `json:"id"`
				Idnumber        string `json:"idnumber"`
				Isfavourite     bool   `json:"isfavourite"`
				Progress        int    `json:"progress"`
				Shortname       string `json:"shortname"`
				Showshortname   bool   `json:"showshortname"`
				Startdate       int    `json:"startdate"`
				Summary         string `json:"summary"`
				Summaryformat   int    `json:"summaryformat"`
				Viewurl         string `json:"viewurl"`
				Visible         bool   `json:"visible"`
			} `json:"course"`
			Deleteurl         string      `json:"deleteurl"`
			Description       string      `json:"description"`
			Descriptionformat int         `json:"descriptionformat"`
			Draggable         bool        `json:"draggable"`
			Editurl           string      `json:"editurl"`
			Eventcount        interface{} `json:"eventcount"`
			Eventtype         string      `json:"eventtype"`
			Formattedtime     string      `json:"formattedtime"`
			Groupid           interface{} `json:"groupid"`
			Groupname         interface{} `json:"groupname"`
			Icon              struct {
				Alttext   string `json:"alttext"`
				Component string `json:"component"`
				Key       string `json:"key"`
			} `json:"icon"`
			ID                      int         `json:"id"`
			Instance                int         `json:"instance"`
			Isactionevent           bool        `json:"isactionevent"`
			Iscategoryevent         bool        `json:"iscategoryevent"`
			Iscourseevent           bool        `json:"iscourseevent"`
			Islastday               bool        `json:"islastday"`
			Location                string      `json:"location"`
			Modulename              string      `json:"modulename"`
			Name                    string      `json:"name"`
			Normalisedeventtype     string      `json:"normalisedeventtype"`
			Normalisedeventtypetext string      `json:"normalisedeventtypetext"`
			Popupname               string      `json:"popupname"`
			Repeatid                interface{} `json:"repeatid"`
			Subscription            struct {
				Displayeventsource bool `json:"displayeventsource"`
			} `json:"subscription"`
			Timeduration    int    `json:"timeduration"`
			Timemodified    int    `json:"timemodified"`
			Timesort        int    `json:"timesort"`
			Timestart       int    `json:"timestart"`
			URL             string `json:"url"`
			Userid          int    `json:"userid"`
			Viewurl         string `json:"viewurl"`
			Visible         int    `json:"visible"`
			Maxdayerror     string `json:"maxdayerror,omitempty"`
			Maxdaytimestamp int    `json:"maxdaytimestamp,omitempty"`
			Mindayerror     string `json:"mindayerror,omitempty"`
			Mindaytimestamp int    `json:"mindaytimestamp,omitempty"`
		} `json:"events"`
		FilterSelector    string `json:"filter_selector"`
		Larrow            string `json:"larrow"`
		Neweventtimestamp int    `json:"neweventtimestamp"`
		Nextperiod        struct {
			Hours     int    `json:"hours"`
			Mday      int    `json:"mday"`
			Minutes   int    `json:"minutes"`
			Mon       int    `json:"mon"`
			Month     string `json:"month"`
			Seconds   int    `json:"seconds"`
			Timestamp int    `json:"timestamp"`
			Wday      int    `json:"wday"`
			Weekday   string `json:"weekday"`
			Yday      int    `json:"yday"`
			Year      int    `json:"year"`
		} `json:"nextperiod"`
		Nextperiodlink string `json:"nextperiodlink"`
		Nextperiodname string `json:"nextperiodname"`
		Periodname     string `json:"periodname"`
		Previousperiod struct {
			Hours     int    `json:"hours"`
			Mday      int    `json:"mday"`
			Minutes   int    `json:"minutes"`
			Mon       int    `json:"mon"`
			Month     string `json:"month"`
			Seconds   int    `json:"seconds"`
			Timestamp int    `json:"timestamp"`
			Wday      int    `json:"wday"`
			Weekday   string `json:"weekday"`
			Yday      int    `json:"yday"`
			Year      int    `json:"year"`
		} `json:"previousperiod"`
		Previousperiodlink string `json:"previousperiodlink"`
		Previousperiodname string `json:"previousperiodname"`
		Rarrow             string `json:"rarrow"`
	} `json:"data"`
	Error bool `json:"error"`
}

func getAttendance(user Users) Attendance {
	now := time.Now()
	var attendance []Attendance
	AttendancePayload := []byte(`[{"args": {"courseid": 1, "day": ` + strconv.Itoa(now.Day()) + `, "month": ` + strconv.Itoa(int(now.Month())) + `, "year": ` + strconv.Itoa(now.Year()) + `}, "index": 0, "methodname": "core_calendar_get_calendar_day_view"}]`)
	req := fasthttp.AcquireRequest()
	req.SetBody(AttendancePayload)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	cookie := strings.Split(user.Cookie, "=")
	req.Header.SetCookie(cookie[0], cookie[1])
	req.SetRequestURI(fmt.Sprint(serviceURL, user.SessKey))
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil{
		panic(err)
	}
	fasthttp.ReleaseRequest(req)
	if err := json.Unmarshal(res.Body(), &attendance); err != nil {
			panic(err)
		}
	fasthttp.ReleaseResponse(res)
	return attendance[0]
}

