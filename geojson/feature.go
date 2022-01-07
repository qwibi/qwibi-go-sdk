package geojson

import (
	"fmt"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

type QGeometry interface {
	// Valid() error
}

//QFeature ...
type QFeature struct {
	Geometry   QGeometry
	Properties *structpb.Struct
}

func (c *QFeature) Valid() error {
	if c.Geometry == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return err
	}

	switch v := c.Geometry.(type) {
	case *QPoint:
		return nil
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return errors.WithStack(err)
	}
}

// NewFeaturePb ...
func NewFeaturePb(pb *proto.QPBxFeature) (*QFeature, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
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
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}

// NewFeature ...
func NewFeature(geometry QGeometry) *QFeature {
	feature := &QFeature{
		Geometry: geometry,
		Properties: &structpb.Struct{
			Fields: make(map[string]*structpb.Value),
		},
	}
	return feature
}

// NewPointFeautre ...
func NewPointFeautre() *QFeature {
	return NewFeature(NewPointZero())
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
		pb.Geometry = &proto.QPBxFeature_Point{
			Point: v.Pb(),
		}
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}

	return pb, nil
}

// SetString ...
func (c *QFeature) SetString(key string, value string) {
	c.Properties.Fields[key] = &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: value,
		},
	}
}

// GetString ...
func (c *QFeature) GetString(key string) string {
	switch v := c.Properties.Fields[key].Kind.(type) {
	case *structpb.Value_StringValue:
		return v.StringValue
	}

	return ""
}
