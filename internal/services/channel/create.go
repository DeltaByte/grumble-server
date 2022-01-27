package channelService

import (
	"context"

	// internal
	pb "github.com/grumblechat/server/gen/go/channel"
	"github.com/grumblechat/server/internal/entities/channel"
	"github.com/grumblechat/server/internal/helpers"

	// external
	"github.com/jinzhu/copier"
)

type createChannelValidator struct {
	Type    pb.ChannelType `validate:"ne=0,required"`
	Name    string         `validate:"max=100,required"`
	Topic   string         `validate:"max=1024"`
	Bitrate uint32         `validate:"min=4,max=255"`
}

func (srv *Server) CreateChannel(ctx context.Context, req *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	// validate request
	if err := helpers.ValidateRequest(&createChannelValidator{}, req); err != nil {
		return &pb.CreateChannelResponse{
			Status: helpers.StatusValidationFailed(err),
		}, nil
	}

	// bind request to new channel
	chn := channel.New(req.GetType())
	copier.Copy(chn, req)

	// persist channel
	if err := channel.Save(srv.db, chn); err != nil {
		return nil, err
	}

	// return created channel to client
	return &pb.CreateChannelResponse{
		Channel: chn,
		Status: helpers.StatusOK(),
	}, nil
}
