package helpers

import (
	"idist-go/app/providers/configProvider"
	"strconv"
	"time"
)

func Now() time.Time {
	timeZone, _ := strconv.Atoi(configProvider.GetConfig().GetString("time.timezone"))
	return time.Now().Add(time.Duration(timeZone) * time.Hour)
}

func ConvertTimeYYYYMMDD(input string) time.Time {
	time, _ := time.Parse("2006-01-02", input)
	return time
}
