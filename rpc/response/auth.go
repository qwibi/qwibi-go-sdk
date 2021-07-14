package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QAuthResponse ...
type QAuthResponse struct {
	Token string
}

func (c *QAuthResponse) Valid() error {
	if c.Token == "" {
		return errors.New("Empty token value")
	}

	return nil
}

// Pb ...
func (c *QAuthResponse) Pb() (*proto.QPBxAuthResponse, error) {

	if err := c.Valid(); err != nil {
		return nil, err
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
		return nil, errors.New("Invalid format")
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
