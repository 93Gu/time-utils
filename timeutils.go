package timeutils

import "time"

// FormatTime formats a time.Time object into a string with the given layout
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime parses a time string into a time.Time object with the given layout
func ParseTime(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}
