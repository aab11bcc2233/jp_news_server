package util

import "time"

func NowTimeStamp() int64  {
	return time.Now().UnixNano() / 1e6
}

func TimeStamp(t time.Time) int64  {
	return time.Time(t).UnixNano() / 1e6
}