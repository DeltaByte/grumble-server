package helpers

import (
	"time"
)

func TouchTimestamp(ts time.Time, now time.Time, nilOnly bool) time.Time {
	if ts.IsZero() || !nilOnly {
		return now
	}
	return ts
}