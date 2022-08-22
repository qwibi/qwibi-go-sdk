package request

// import (
// 	"github.com/pkg/errors"
// 	"github.com/qwibi/qwibi-go-sdk/proto"
// 	"github.com/rs/zerolog/log"
// )

// // QStreamRequest ...
// type QStreamRequest struct {
// 	Request *proto.QPBxStreamRequest
// }

// func (c *QStreamRequest) Valid() error {
// 	if c.Request == nil {
// 		err := errors.New("Stream request is not defined")
// 		log.Error().Stack().Err(err).Msg("")
// 		return errors.WithStack(err)
// 	}

// 	return nil
// }

// // Pb ...
// func (c *QStreamRequest) Pb() (*proto.QPBxStreamRequest, error) {
// 	if err := c.Valid(); err != nil {
// 		return nil, err
// 	}

// 	return c.Request, nil
// }

// // NewStreamRequestPb ...
// func NewStreamRequestPb(pb *proto.QPBxStreamRequest) (*QStreamRequest, error) {
// 	if pb == nil {
// 		err := errors.New("Invalid parameter type nil")
// 		log.Error().Stack().Err(err).Msg("")
// 		return nil, errors.WithStack(err)
// 	}

// 	req := &QStreamRequest{
// 		Request: pb,
// 	}

// 	return req, nil
// }

// // NewStreamRequest ...
// func NewStreamRequest() (*QStreamRequest, error) {
// 	req := &QStreamRequest{
// 		Request: &proto.QPBxStreamRequest{},
// 	}

// 	return req, nil
// }

// // NewStreamRequest ...
// // func NewStreamRequest(m interface{}) (*QStreamRequest, error) {

// // 	var pb *proto.QPBxStreamRequest

// // 	switch v := m.(type) {
// // 	case *:
// // 		pbFeature, err := v.Pb()
// // 		if err != nil {
// // 			return nil, errors.WithStack(err)
// // 		}
// // 		postRequest := &proto.QPBxPostRequest{Feature: pbFeature}
// // 		postStreamRequest := &proto.QPBxStreamRequest_Post{Post: postRequest}
// // 		pb = &proto.QPBxStreamRequest{Request: postStreamRequest}
// // 	default:
// // 		err := errors.New("Unknown feature type")
// // 		log.Error().Stack().Err(err).Msg("")
// // 		return nil, errors.WithStack(err)
// // 	}

// // 	r := &QStreamRequest{
// // 		Request: pb,
// // 	}

// // 	return r, nil
// // }
