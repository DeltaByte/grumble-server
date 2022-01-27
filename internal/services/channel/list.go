package channelService

import (
	"context"

	pb "github.com/grumblechat/server/gen/go/channel"
	"github.com/grumblechat/server/internal/entities/channel"
)

func (srv *Server) ListChannels(ctx context.Context, in *pb.ListChannelsRequest) (*pb.ListChannelsResponse, error) {
	// list all channels
	channels, err := channel.GetAll(srv.db)

	// handle errors
	if err != nil {
		return nil, err
	}

	// respond to client
	return &pb.ListChannelsResponse{
		Channels: channels,
	}, nil
}