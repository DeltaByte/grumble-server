package channel

import (
	"bytes"
	"encoding/gob"
	"errors"

	bolt "go.etcd.io/bbolt"
)

const BoltBucketName = "channels"

type Channel interface {
	GetType() string
	Encode() ([]byte, error)
	Decode([]byte) error
	Save(*bolt.DB) error
}

func Decode(encoded []byte) (Channel, error) {
	// extract type
	type baseChannelType struct { Type string }
	baseChannel := &baseChannelType{}
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(baseChannel)
	if err != nil { return nil, err }

	// decode VoiceChannel
	if (baseChannel.Type == "voice") {
		vc := &VoiceChannel{}
		vc.Decode(encoded)
		return vc, nil
	}

	// decode TextChannel
	if (baseChannel.Type == "text") {
		tc := &TextChannel{}
		tc.Decode(encoded)
		return tc, nil
	}

	// handle unknown types
	return nil, errors.New("unknown channel type")
}