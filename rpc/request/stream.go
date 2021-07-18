package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QStreamRequest ...
type QStreamRequest struct {
	Request *proto.QPBxStreamRequest
}

func (c *QStreamRequest) Valid() error {
	if c.Request == nil {
		return errors.New("Invalid request type nil")
	}

	return nil
}

// Pb ...
func (c *QStreamRequest) Pb() (*proto.QPBxStreamRequest, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return c.Request, nil
}

// NewStreamRequestPb ...
func NewStreamRequestPb(pb *proto.QPBxStreamRequest) (*QStreamRequest, error) {
	if pb == nil {
		return nil, errors.New("Invalid format type nil")
	}

	req := &QStreamRequest{
		Request: pb,
	}

	return req, nil
}

// NewStreamRequest ...
func NewStreamRequest() (*QStreamRequest, error) {

	var pb *proto.QPBxStreamRequest

	r := &QStreamRequest{
		Request: pb,
	}

	return r, nil
}
