package geometry

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QGeometry ...
type QGeometry interface {
	Valid() error
	Pb() (*proto.QPBxGeometry, error)
}

// NewGeometryPb ...
func NewGeometryPb(pb *proto.QPBxGeometry) (QGeometry, error) {
	if pb == nil {
		err := errors.New("Invalid geometry type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	switch v := pb.Geometry.(type) {
	case *proto.QPBxGeometry_Layer:
		return NewLayerPb(v.Layer)
	case *proto.QPBxGeometry_Point:
		return NewPointPb(v.Point)
	default:
		err := errors.New("Unknown geometry type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}
