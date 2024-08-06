package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QQueryResponse struct {
	RequestId string
	Objects   []*geo.QGeoObject
}

func NewQueryResponse(requestId string, objects []*geo.QGeoObject) (*QQueryResponse, error) {
	res := &QQueryResponse{
		RequestId: requestId,
		Objects:   objects,
	}
	return res, nil
}

func NewQueryResponsePb(in *proto.QPBxQueryResponse) (*QQueryResponse, error) {
	objects := make([]*geo.QGeoObject, len(in.Objects))
	for _, pbObject := range in.Objects {
		object, err := geo.NewGeoObjectPb(pbObject)
		if err != nil {
			return nil, qlog.Error(err)
		}
		objects = append(objects, object)
	}

	return NewQueryResponse(in.RequestId, objects)
}

func (c *QQueryResponse) Pb() *proto.QPBxQueryResponse {
	pbObjects := make([]*proto.QPBxGeoObject, len(c.Objects))
	for _, object := range c.Objects {
		pbObjects = append(pbObjects, object.Pb())
	}

	pb := &proto.QPBxQueryResponse{
		RequestId: c.RequestId,
		Objects:   pbObjects,
	}

	return pb
}
