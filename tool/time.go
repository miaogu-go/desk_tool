package tool

import "time"

func GetCurrentUnixTime() int64 {
	return time.Now().Unix()
}
