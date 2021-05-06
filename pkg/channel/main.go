package channel

import "github.com/dgraph-io/badger/v3"

type Channel interface {
	Encode() ([]byte, error)
	Save(*badger.DB) error
}