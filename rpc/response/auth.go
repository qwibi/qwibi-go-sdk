package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QAuthResponse ...
type QAuthResponse struct {
	Token string
}

func (c *QAuthResponse) Valid() error {
	if c.Token == "" {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QAuthResponse) Pb() (*proto.QPBxAuthResponse, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	token := "TODO"

	pb := &proto.QPBxAuthResponse{
		Token: token,
	}

	return pb, nil
}

// NewAuthResponsePb ...
func NewAuthResponsePb(pb *proto.QPBxAuthResponse) (*QAuthResponse, error) {
	if pb == nil {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	res := &QAuthResponse{
		Token: pb.Token,
	}

	if err := res.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

// NewAuthResponse ...
func NewAuthResponse(token string) (*QAuthResponse, error) {

	r := &QAuthResponse{
		Token: token,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
