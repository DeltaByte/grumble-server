package channelService

import (
	pbService "github.com/grumblechat/server/gen/go/channel"

	bolt "go.etcd.io/bbolt"
)

type Server struct {
	pbService.UnimplementedChannelServiceServer
	db *bolt.DB
}

func New(db *bolt.DB) *Server {
	return &Server{
		db: db,
	}
}