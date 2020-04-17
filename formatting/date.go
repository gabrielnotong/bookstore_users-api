package formatting

import "time"

const dateLayout = "2006-01-02 15:03:05"

func DateNow() time.Time {
	return time.Now().UTC()
}

func DateNowString() string {
	return DateNow().Format(dateLayout)
}
