package timeutil

import "time"

const timeFormat = "2006-01-02 15:04:05"

// TimeFormat format time like "2006-01-02 15:04:05"
func TimeFormat(t time.Time) string {
	return t.Format(timeFormat)
}

// ZeroTime return time.Time = "1970-01-01 08:00:00"
func ZeroTime() time.Time {
	return time.Unix(0, 0)
}
