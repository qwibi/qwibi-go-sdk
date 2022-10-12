package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
	"github.com/qwibi/qwibi-go-sdk/utils"
)

type QGeoLayer struct {
	Gid        string
	Properties []byte
}

func NewGeoLayer(gid string) *QGeoLayer {
	if gid == "" {
		gid = utils.NewID()
	}
	return &QGeoLayer{
		Gid: gid,
	}
}

func NewGeoLayerPb(pb *proto.QPBxGeoObject) (*QGeoLayer, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	layerPb := pb.GetLayer()
	if layerPb == nil {
		err := errors.New("Invalid GeoLayer format type nil")
		return nil, qlog.Error(err)
	}

	layer := &QGeoLayer{
		Gid:        layerPb.Gid,
		Properties: layerPb.Properties,
	}

	return layer, layer.Valid()
}

func (c *QGeoLayer) Valid() error {
	return nil
}

func (c *QGeoLayer) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Geo: &proto.QPBxGeoObject_Layer{
			Layer: &proto.QPBxGeoLayer{
				Gid:        c.Gid,
				Properties: c.Properties,
			},
		},
	}
}
