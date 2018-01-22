package util

import "time"

func GetTimeString(from string) string {
	parsed, err := time.Parse("2006-01-02T15:04:05Z", from)

	if err != nil {
		return from
	}

	return parsed.Format("2006-01-02 15:04:05")
}
