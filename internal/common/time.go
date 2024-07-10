package common

import "time"

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}
