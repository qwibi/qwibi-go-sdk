package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QJoinResponse struct {
	RequestId string
	Layer     *layer.QLayer
	Objects   []*geo.QGeoObject
}

func NewJoinResponse(layer *layer.QLayer, objects []*geo.QGeoObject) (*QJoinResponse, error) {
	res := &QJoinResponse{
		RequestId: utils.RequestId(),
		Layer:     layer,
		Objects:   objects,
	}

	return res, res.Valid()
}

func NewJoinResponsePb(in *proto.QPBxJoinResponse) (*QJoinResponse, error) {
	layer, err := layer.NewGeoLayerPb(in.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	objects := make([]*geo.QGeoObject, 0)
	for _, pbObject := range in.Objects {
		object, err := geo.NewGeoObjectPb(pbObject)
		if err != nil {
			return nil, qlog.Error(err)
		}
		objects = append(objects, object)
	}

	return NewJoinResponse(layer, objects)
}

func (c *QJoinResponse) Pb() *proto.QPBxJoinResponse {
	objects := make([]*proto.QPBxGeoObject, 0)
	for _, object := range c.Objects {
		if object != nil {
			objects = append(objects, object.Pb())
		}
	}

	return &proto.QPBxJoinResponse{
		RequestId: c.RequestId,
		Layer:     c.Layer.Pb(),
		Objects:   objects,
	}
}

func (c *QJoinResponse) Valid() error {
	if c.RequestId == "" {
		return qlog.Error("request ID is not defined")
	}

	if c.Layer == nil {
		return qlog.Error("layer is not defined")
	}

	if c.Objects == nil {
		return qlog.Error("bad parameter type nil")
	}

	return nil
}
