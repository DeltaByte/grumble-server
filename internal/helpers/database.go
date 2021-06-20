package helpers

import (
	"time"

	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TouchTimestamp(pbTS *timestamppb.Timestamp, now time.Time, nilOnly bool) {
	ts := pbTS.AsTime()

	if ts.IsZero() || !nilOnly {
		ts = now
	}

	pbTS.Seconds = ts.Unix()
	pbTS.Nanos = int32(ts.Nanosecond())
}

func ParseKSUIDBytes(id string) ([]byte, error) {
	kid, err := ksuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return kid.Bytes(), nil
}
