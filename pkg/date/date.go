package date

import (
	"log"
	"time"
)

const (
	DATE_PATTERN         = "2006-01-02T15:04:05"
	DATE_PATTERN_WITH_TZ = "2006-01-02T15:04:05 MST"
)

func Format(date time.Time, timezone string) string {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Failed to get timezone for %s, error: %s", timezone, err)
		return ""
	}
	return date.In(location).Format(DATE_PATTERN)
}

func FormatWithTZ(date time.Time, timezone string) string {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Failed to get timezone for %s, error: %s", timezone, err)
		return ""
	}
	return date.In(location).Format(DATE_PATTERN_WITH_TZ)
}

// func main() {
//     now := time.Now()
//
//     fmt.Println(now)
//     fmt.Println(now.Format(time.RFC3339))
//     fmt.Println(now.Format("2006-01-02T15:04:05.000 MST -07:00"))
//     fmt.Println(now.Local().Format("2006-01-02T15:04:05.000 MST -07:00"))
//     location, _ := time.LoadLocation("Asia/Shanghai")
//     fmt.Println(now.In(location).Format("2006-01-02T15:04:05.000 MST -07:00"))
//
//     utc, _ := time.LoadLocation("")
//     fmt.Println(now.In(utc).Format("2006-01-02T15:04:05.000 MST -07:00"))
//
//     est, _ := time.LoadLocation("EST")
//     fmt.Println(now.In(est).Format("2006-01-02T15:04:05.000 MST -07:00"))
//     // https://golangbyexample.com/time-date-formatting-in-go/
// }
