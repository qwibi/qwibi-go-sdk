package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geometry"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QGeoObject ...
type QGeoObject struct {
	Gid        string
	Geometry   geometry.QGeometry
	Properties *structpb.Struct
}

// NewGeoObjectPb ...
func NewGeoObjectPb(pb *proto.QPBxGeoObject) (*QGeoObject, error) {
	if pb == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	if pb.Gid == "" {
		pb.Gid = utils.NewID()
	}

	if pb.Geometry == nil {
		err := errors.New("Invalid feature type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	geometry, err := geometry.NewGeometryPb(pb.Geometry)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object := &QGeoObject{
		Gid:        pb.Gid,
		Geometry:   geometry,
		Properties: pb.Properties,
	}

	return object, nil
}

// Valid ...
func (c QGeoObject) Valid() error {
	if c.Gid == "" {
		err := errors.New("Object gid is not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if c.Geometry == nil {
		err := errors.New("Object geometry is not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return c.Geometry.Valid()
}

// Pb ...
func (c QGeoObject) Pb() (*proto.QPBxGeoObject, error) {
	pb := &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Properties: c.Properties,
	}

	geometryPb, err := c.Geometry.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb.Geometry = geometryPb

	return pb, nil
}
