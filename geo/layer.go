package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
)

type QGeoLayer struct {
	Gid        string
	Properties []byte
}

func NewGeoLayer() *QGeoLayer {
	return &QGeoLayer{
		Gid: utils.NewID(),
	}
}

func NewGeoLayerPb(pb *proto.QPBxGeoObject) (*QGeoLayer, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	layerPb := pb.GetLayer()
	if layerPb == nil {
		err := errors.New("Invalid GeoLayer format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
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
