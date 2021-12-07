package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)

type Assignment struct {
	Data struct {
		Events []struct {
			Action struct {
				Actionable    bool   `json:"actionable"`
				Itemcount     int    `json:"itemcount"`
				Name          string `json:"name"`
				Showitemcount bool   `json:"showitemcount"`
				URL           string `json:"url"`
			} `json:"action"`
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
			Location                string      `json:"location"`
			Modulename              string      `json:"modulename"`
			Name                    string      `json:"name"`
			Normalisedeventtype     string      `json:"normalisedeventtype"`
			Normalisedeventtypetext string      `json:"normalisedeventtypetext"`
			Repeatid                interface{} `json:"repeatid"`
			Subscription            struct {
				Displayeventsource bool `json:"displayeventsource"`
			} `json:"subscription"`
			Timeduration int    `json:"timeduration"`
			Timemodified int    `json:"timemodified"`
			Timesort     int    `json:"timesort"`
			Timestart    int    `json:"timestart"`
			URL          string `json:"url"`
			Userid       int    `json:"userid"`
			Viewurl      string `json:"viewurl"`
			Visible      int    `json:"visible"`
		} `json:"events"`
		Firstid int `json:"firstid"`
		Lastid  int `json:"lastid"`
	} `json:"data"`
	Error bool `json:"error"`
}
func getAssignments(user Users) Assignment{
	now := time.Now()
	AssignmentPayload := []byte(`[{"index":0,"methodname":"core_calendar_get_action_events_by_timesort","args":{"limitnum":26,"timesortfrom":` + strconv.FormatInt(now.Unix(), 10) + `,"timesortto":` + strconv.FormatInt(now.Unix()+15482362, 10) + `,"limittononsuspendedevents":true}}]`)
	var assignment []Assignment
	req := fasthttp.AcquireRequest()
	req.SetBody(AssignmentPayload)
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
	if err := json.Unmarshal(res.Body(), &assignment); err != nil {
			panic(err)
		}
	fasthttp.ReleaseResponse(res)
	return assignment[0]
}
