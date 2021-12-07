package main

import (
	"fmt"
	"github.com/icza/gox/timex"
	"time"
)

type Users struct {
	ID uint `json:"id"`
	NIM string `json:"nim"`
	Name string `json:"name"`
	Pass string `json:"pass"`
	SessKey string `json:"sess_key"`
	Cookie string `json:"cookie"`
}

type Deadline struct {
	Name string
	Subject string
	Due string
}

func parseEpoch(futureEpoch int) string {
	_, month, day, hour, min, sec := timex.Diff(time.Unix(int64(futureEpoch), 0), time.Now())
	var duration string
	if int64(futureEpoch) > time.Now().Unix() {
		duration = fmt.Sprintf("Due in %d months, %d days, %d hours, %d mins and %d seconds.",
    month, day, hour, min, sec)
	} else {
		duration = fmt.Sprintf("Late %d months, %d days, %d hours, %d mins and %d seconds.",
    month, day, hour, min, sec)
	}
	return duration
}