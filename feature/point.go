package feature

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geometry"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

type QPointFeature struct {
	Geometry   *geometry.QPoint
	Properties []byte
}

func NewPointFeature() *QPointFeature {
	return &QPointFeature{
		Geometry: geometry.NewPoint(),
	}
}

func NewPointFeaturePb(pb *proto.QPBxPointFeature) (*QPointFeature, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	geometry, err := geometry.NewPointPb(pb.Geometry)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	feature := &QPointFeature{
		Geometry:   geometry,
		Properties: pb.Properties,
	}

	return feature, feature.Valid()
}

func (c *QPointFeature) Valid() error {
	if c.Geometry == nil {
		err := errors.New("Invalid geometry type nil")
		log.Error().Stack().Err(err).Msg("")
		return err
	}

	return c.Geometry.Valid()
}

func (c *QPointFeature) Pb() *proto.QPBxPointFeature {
	return &proto.QPBxPointFeature{
		Geometry:   c.Geometry.Pb(),
		Properties: c.Properties,
	}
}
