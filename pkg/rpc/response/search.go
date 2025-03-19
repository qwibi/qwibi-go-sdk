package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QSearchResponse struct {
	RequestId string
	Layers    []*layer.QLayer
}

func NewSearchResponse(requestId string, layers []*layer.QLayer) (*QSearchResponse, error) {
	res := &QSearchResponse{
		RequestId: requestId,
		Layers:    layers,
	}
	return res, nil
}

func NewSearchResponsePb(in *proto.QPBxSearchResponse) (*QSearchResponse, error) {
	layers := make([]*layer.QLayer, 0)
	for _, pb := range in.Layers {
		layer, err := layer.NewLayerPb(pb)
		if err != nil {
			return nil, qlog.Error(err)
		}
		layers = append(layers, layer)
	}

	return NewSearchResponse(in.RequestId, layers)
}

func (c *QSearchResponse) Pb() *proto.QPBxSearchResponse {
	pbLayers := make([]*proto.QPBxLayer, 0)
	for _, layer := range c.Layers {
		pbLayers = append(pbLayers, layer.Pb())
	}

	pb := &proto.QPBxSearchResponse{
		RequestId: c.RequestId,
		Layers:    pbLayers,
	}

	return pb
}
