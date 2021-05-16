package pagination

import (
	"bytes"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type Pagination struct {
	Cursor     ksuid.KSUID `query:"cursor"`
	NextCursor ksuid.KSUID
	Count      uint16      `query:"count" validate:"min=1,max=1000"`
	Reverse    bool        `query:"reverse"`
}

func (pgn *Pagination) InitCursor(cursor *bolt.Cursor) (key []byte, value []byte) {
	if (!pgn.Cursor.IsNil()) {
		// seek to the specified key, Bolt automatically goes to the next if it isn't found
		seekKey := pgn.Cursor.Bytes()
		k, v := cursor.Seek(seekKey)

		// manually go to the next key if the found one is the same as the pagination cursor
		if (bytes.Equal(seekKey, k)) {
			k, v = pgn.MoveCursor(cursor)
		}

		return k, v
	}

	// start from end of bucket if paginating in reverse order
	if (pgn.Reverse) {
		return cursor.Last()
	}

	return cursor.First()
}

func (pgn *Pagination) MoveCursor(cursor *bolt.Cursor) (key []byte, value []byte) {
	if (pgn.Reverse) {
		return cursor.Prev()
	}

	return cursor.Next()
}

func (pgn *Pagination) EndKey(cursor *bolt.Cursor) (ksuid.KSUID, error) {
	key, _ := cursor.Prev()

	if (pgn.Reverse) {
		key, _ = cursor.Next()
	}

	if (key == nil) {
		return ksuid.Nil, nil
	}
	
	return ksuid.FromBytes(key)
}