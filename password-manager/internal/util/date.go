package util

import "time"

func DateTimeNow() int64 {
	currentTime := time.Now()
	return currentTime.Unix()
}
