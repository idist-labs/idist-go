package helpers

import (
	"ai-camera-api-cms/app/providers/configProvider"
	"strconv"
	"time"
)

func Now() time.Time {
	timeZone, _ := strconv.Atoi(configProvider.GetConfig().GetString("time.timezone"))
	return time.Now().Add(time.Duration(timeZone) * time.Hour)
}

func PNow() *time.Time {
	timeZone, _ := strconv.Atoi(configProvider.GetConfig().GetString("time.timezone"))
	now := time.Now().Add(time.Duration(timeZone) * time.Hour)
	return &now
}
