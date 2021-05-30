package message

import (
	"bytes"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/grumblechat/server/internal/database"
	"github.com/grumblechat/server/internal/entities/channel"
	"github.com/grumblechat/server/internal/entities/message"
)

func TestMessageOrdering(t *testing.T) {
	t.Parallel()
	var waitGroup sync.WaitGroup

	// default settings
	chunkSize := 1000
	iter := 1000000

	// make test shorter if needed
	if testing.Short() {
		iter = 1000
		chunkSize = 100
		t.Logf("Running in short mode (%d messages)", iter)
	}

	// setup
	dbDir := t.TempDir()
	db := database.Init(dbDir)
	defer db.Close()

	// create channel
	chn := channel.NewText()
	chn.Name = t.Name()
	chn.Save(db)

	// setup channel for message generation
	msgChan := make(chan []*message.Message)
	var msgs []*message.Message

	// batch message generation into chunks
	chunks := iter / chunkSize
	waitGroup.Add(chunks)
	for chunk := 0; chunk < chunks; chunk++ {
		go func() {
			defer waitGroup.Done()
			var msgChunk []*message.Message

			// generate batch of 100 messages
			for i := 0; i < chunkSize; i++ {
				msg := message.New(chn.ID)
				msg.Body = "Lorem Ipsum"
				msg.CreatedAt = time.Now()
				msgChunk = append(msgChunk, msg)
			}

			// save messages as batch
			if err := message.BatchSave(db, msgChunk); err != nil {
				t.Error(err)
			}

			// write messages back to channel
			msgChan <- msgChunk
		}()
	}

	// wait for message generation and close channel
	go func() {
		waitGroup.Wait()
		close(msgChan)
	}()

	// write messages from channel in to slice,
	// this is blocked if the channel is not closed
	for msg := range msgChan {
		msgs = append(msgs, msg...)
	}

	// sort by KSUID
	waitGroup.Add(1)
	var sortedById []*message.Message
	go func() {
		toSort := msgs
		sort.Slice(toSort, func(i, j int) bool {
			id1 := toSort[i].ID.Bytes()
			id2 := toSort[j].ID.Bytes()
			return bytes.Compare(id1, id2) < 0
		})
		sortedById = toSort
		waitGroup.Done()
	}()

	// sort by time
	waitGroup.Add(1)
	var sortedByTime []*message.Message
	go func() {
		toSort := msgs
		sort.Slice(toSort, func(i, j int) bool {
			t1 := toSort[i].CreatedAt
			t2 := toSort[j].CreatedAt
			return t1.Before(t2)
		})
		sortedByTime = toSort
		waitGroup.Done()
	}()

	// wait for sorting to complete
	waitGroup.Wait()

	// see if the lists are the same
	correct := 0
	for i := 0; i < iter; i++ {
		byId := sortedById[i]
		byTime := sortedByTime[i]
		same := byId.ID == byTime.ID
		if same {
			correct++
		}
	}

	// TODO: also iterate through DB to ensure ordering consistency

	// check results
	delta := (float64(correct) * float64(100)) / float64(iter)
	if correct != iter {
		t.Errorf("Correct: %d, Incorrect %d, Delta: %.3f%%, ", correct, (iter - correct), delta)
	}
}
