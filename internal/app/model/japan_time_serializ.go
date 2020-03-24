package model

import (
	"fmt"
	"jpnews/internal/pkg/util"
	"time"
)

type Time time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	timeStr := string(data)
	now, err := time.ParseInLocation("\""+timeFormat+"\"", timeStr, time.FixedZone("UTC", 9*3600))
	if err != nil {
		return err
	}
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	//do your serializing here
	//stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	ts := util.TimeStamp(time.Time(t))
	stamp := fmt.Sprintf("%d", ts)
	return []byte(stamp), nil
}
