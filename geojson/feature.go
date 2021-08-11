package geojson

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

//QFeature ...
type QFeature struct {
	Gid        string
	Geometry   QGeometry
	Properties *structpb.Struct
}

func (c *QFeature) Valid() error {
	if c.Gid == "" || c.Geometry == nil || c.Geometry.Valid() != nil {
		return errors.New("QFeature invalid format")
	}

	return nil
}

// NewFeaturePb ...
func NewFeaturePb(pb *proto.QPBxFeature) (*QFeature, error) {
	if pb == nil {
		return nil, errors.New("Invalid format type nil")
	}

	feature := &QFeature{
		Gid:        pb.Gid,
		Properties: pb.Properties,
	}

	switch v := pb.Geometry.Object.(type) {
	case *proto.QPBxGeometry_Point:
		point, err := NewPointPb(v.Point)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		feature.Geometry = point
		return feature, nil
	default:
		return nil, errors.New("Unknown geojson object")
	}
}

// NewFeature ...
func NewFeature(gid string, geometry QGeometry) *QFeature {
	feature := &QFeature{
		Gid:      gid,
		Geometry: geometry,
		Properties: &structpb.Struct{
			Fields: make(map[string]*structpb.Value),
		},
	}
	return feature
}

// NewPointFeautre ...
func NewPointFeautre(gid string) *QFeature {
	return NewFeature(gid, NewPointZero())
}

// Pb ...
func (c *QFeature) Pb() (*proto.QPBxFeature, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pbGeometry, err := c.Geometry.GeometryPb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxFeature{
		Gid:        c.Gid,
		Geometry:   pbGeometry,
		Properties: c.Properties,
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
