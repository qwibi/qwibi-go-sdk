package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QStreamResponse ...
type QStreamResponse struct {
	Response *proto.QPBxStreamResponse
}

func (c *QStreamResponse) Valid() error {
	if c.Response == nil {
		return errors.New("Invalid response type nil")
	}

	return nil
}

// Pb ...
func (c *QStreamResponse) Pb() (*proto.QPBxStreamResponse, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := c.Response

	return pb, nil
}

//// NewStreamResponsePb ...
//func NewStreamResponsePb(pb *proto.QPBxStreamResponse) (*QStreamResponse, error) {
//	if pb == nil {
//		return nil, errors.New("Invalid format type nil")
//	}
//
//	switch v := pb.Response.(type) {
//	case *proto.QPBxStreamResponse_Post:
//		req, err := NewPostResponsePb(v.Post)
//		if err != nil {
//			logrus.Error(err)
//			return nil, err
//		}
//		logrus.Info("TODO", req, err)
//	default:
//		err := errors.Error_UNKNOWN(v)
//		logrus.Error(err)
//		return nil, err
//	}
//
//	err := fmt.Errorf("TODO")
//	logrus.Error(err)
//
//	return nil, err
//}
//
//// NewStreamResponse ...
//func NewStreamResponse(m interface{}) (*QStreamResponse, error) {
//	switch v := m.(type) {
//	case *QPostResponse:
//		postPb, err := v.Pb()
//		if err != nil {
//			logrus.Error(err)
//			return nil, err
//		}
//		resPb := &proto.QPBxStreamResponse{
//			Response: &proto.QPBxStreamResponse_Post{
//				Post: postPb,
//			},
//		}
//		res := &QStreamResponse{
//			Response: resPb,
//		}
//		return res, nil
//	case *request.QDropRequest:
//		reqPb, err := v.Pb()
//		if err != nil {
//			logrus.Error(err)
//			return nil, err
//		}
//		resPb := &proto.QPBxStreamResponse{
//			Response: &proto.QPBxStreamResponse_Drop{
//				Drop: reqPb,
//			},
//		}
//		res := &QStreamResponse{
//			Response: resPb,
//		}
//		return res, nil
//	default:
//		err := errors.Error_UNKNOWN(v)
//		logrus.Error(err)
//		return nil, err
//	}
//}


/*
// QStreamResponse ...
type QStreamResponse interface {
	Pb() (*proto.QPBxStreamResponse, error)
}

/*
// NewMessage ...
func NewMessage(object interface{}) *QStreamResponse {
	m := &QStreamResponse{
		Object: object,
	}

	return m
}

// Pb ...
func (c *QStreamResponse) Pb() (*proto.QPBxStreamResponse, error) {
	pb := &proto.QPBxStreamResponse{}
	switch v := c.Object.(type) {
	case *request.QPostRequest:
		pb.Object = &proto.QPBxStreamResponse_Post{
			Post: v.Pb(),
		}
	default:
		return nil, util.ErrorUnknown(v)
	}
	return pb, nil
}
*/
