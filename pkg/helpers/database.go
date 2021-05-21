package helpers

import (
	"time"
)

func TouchTimestamp(ts time.Time, nilOnly bool) time.Time {
	if ts.IsZero() || !nilOnly {
		return time.Now()
	}
	return ts
}