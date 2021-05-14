package channel

import bolt "go.etcd.io/bbolt"

type Channel interface {
	Encode() ([]byte, error)
	Save(*bolt.DB) error
}