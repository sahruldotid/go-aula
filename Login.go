package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
)

func login(user Users) (string, []*http.Cookie) {

	var sesskey string
	var logintoken string
	var cookie []*http.Cookie

	// Retrieve Login Token
	loginCollector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"),
		)
	loginCollector.AllowURLRevisit = true
	loginCollector.OnHTML("input[name=logintoken]:nth-of-type(1)", func(e *colly.HTMLElement) {
		logintoken = e.Attr("value")
		credential := map[string]string{"username": user.NIM, "password": user.Pass, "logintoken": logintoken}
		loginCollector.Post(loginURL, credential)
	})
	loginCollector.Visit(loginURL)
	//

	// Retrive Session Key
	sessCollector := loginCollector.Clone()
	sessCollector.OnRequest(func(r *colly.Request) {
		cookie = sessCollector.Cookies(r.URL.String())
	})
	sessCollector.OnHTML("input[name=sesskey]:nth-of-type(1)", func(h *colly.HTMLElement) {
		sesskey = h.Attr("value")
	})
	sessCollector.Visit(sessURL)
	//
	return sesskey, cookie
}

func (user *Users) isLoggedIn() bool {
	return true
}

func checkSession(user Users) bool {
	if (user.Cookie != "" && user.SessKey != "") {
		req := fasthttp.AcquireRequest()
		req.SetBody([]byte(`[{"index":0,"methodname":"core_course_get_recent_courses","args":{"userid":"1","limit":10}}]`))
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
		if !(strings.Contains(string(res.Body()), "Your session has most likely timed out") || strings.Contains(string(res.Body()), "Web service is not available")){
			return true
		}
		fasthttp.ReleaseResponse(res)
	}
	return false
}
