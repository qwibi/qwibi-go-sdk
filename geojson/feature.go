package geojson

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QFeature ...
type QFeature struct {
	Geometry   QGeometry
	Properties *structpb.Struct
}

// NewPointFeautre ...
func NewPointFeautre() *QFeature {
	feature := &QFeature{
		Geometry:   NewZeroPoint(),
		Properties: &structpb.Struct{},
	}

	return feature
}

// NewFeaturePb ...
func NewFeaturePb(pb *proto.QPBxFeature) (*QFeature, error) {
	if pb == nil {
		err := errors.New("Invalid feature type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	feature := &QFeature{
		Properties: pb.Properties,
	}

	switch v := pb.Geometry.(type) {
	case *proto.QPBxFeature_Point:
		point, err := NewPointPb(v.Point)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		feature.Geometry = point
		return feature, nil
	default:
		err := errors.New("Unknown geometry type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}

// Valid
func (c *QFeature) Valid() error {
	return c.Geometry.Valid()
}

// Pb ...
func (c *QFeature) Pb() (*proto.QPBxFeature, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxFeature{
		Properties: c.Properties,
	}

	geometry := c.Geometry
	switch v := geometry.(type) {
	case *QPoint:
		pointPb, err := v.Pb()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		pb.Geometry = &proto.QPBxFeature_Point{
			Point: pointPb,
		}
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}

	return pb, nil
}
