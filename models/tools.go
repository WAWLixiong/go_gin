package models

import "time"

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
	return t
}
