package message_test

import (
	"testing"

	"github.com/grumblechat/server/internal/database"
	"github.com/grumblechat/server/internal/entities/channel"
	"github.com/grumblechat/server/internal/entities/message"
	"github.com/grumblechat/server/test"

	"github.com/google/go-cmp/cmp"
	"github.com/segmentio/ksuid"
)

func TestEncodeDecode(t *testing.T) {
	t.Parallel()

	chnID := ksuid.New()

	// create and persist message
	msg := message.New(chnID)
	msg.Body = test.RandomString(300, test.CharsetDefault)

	// encode message
	enc, err := msg.Encode()
	if err != nil {
		t.Fatal(err)
	}

	// decode message
	dec, err := message.Decode(enc)
	if err != nil {
		t.Fatal(err)
	}

	// compare encoded/decoded
	if !cmp.Equal(dec, msg) {
		t.Error("encoded/decoded messages do not match")
	}
}

func TestSaveFind(t *testing.T) {
	t.Parallel()

	// setup
	dbDir := t.TempDir()
	db := database.Init(dbDir)
	defer db.Close()

	// create channel
	chn := channel.NewText()
	chn.Name = t.Name()
	chn.Save(db)

	// create and persist message
	msg := message.New(chn.ID)
	msg.Body = test.RandomString(300, test.CharsetDefault)
	msg.Save(db)

	// retrieve from DB
	dbMsg, err := message.Find(db, chn.ID, msg.ID)
	if err != nil {
		t.Fatal(err)
	}

	// compare messages
	if dbMsg == nil || !cmp.Equal(msg, dbMsg) {
		t.Error("Failed to retrieve message intact from DB")
	}
}
