package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QJoinRequest ...
type QMessageRequest struct {
	Message string
}

func (c *QMessageRequest) Valid() error {
	if c.Message == "" {
		return errors.New("LayerID is not defined")
	}

	return nil
}

// Pb ...
func (c *QMessageRequest) Pb() (*proto.QPBxMessageRequest, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxMessageRequest{
		Message: c.Message,
	}

	return pb, nil
}

func NewMessageRequest(message string) (*QMessageRequest, error) {
	r := &QMessageRequest{
		Message: message,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
