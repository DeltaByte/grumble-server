package database

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

var (
	channelsDB badger.DB
)

func Init(path string) {
	// channels
	channelsDB, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("%s/channels.db", path)))
	if err != nil { panic("Failed to open database: channels") }
	defer channelsDB.Close()
}

func DBManager(name string) *badger.DB {
	if (name == "channels") {
		return &channelsDB
	}

	panic(fmt.Sprintf("Tried to open unknown database: %s", name))
}
