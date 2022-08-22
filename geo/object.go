package geo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

type QGeoObject interface {
	Valid() error
	Pb() *proto.QPBxGeoObject
}

// NewGeoObjectPb ...
func NewGeoObjectPb(pb *proto.QPBxGeoObject) (QGeoObject, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	switch v := pb.Geo.(type) {
	case *proto.QPBxGeoObject_Layer:
		return NewGeoLayerPb(pb)
	case *proto.QPBxGeoObject_Point:
		return NewGeoPointPb(pb)
	default:
		err := errors.New("Unknown geometry type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}
