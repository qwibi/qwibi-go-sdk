package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geojson"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QGeoObject ...
type QGeoObject struct {
	Gid        string
	Feature    *geojson.QFeature
	Properties *structpb.Struct
}

// NewGeoPoint ...
func NewGeoPoint() *QGeoObject {
	point := &QGeoObject{
		Gid:        utils.NewID(),
		Feature:    geojson.NewPointFeautre(),
		Properties: &structpb.Struct{},
	}
	return point
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

	if pb.Feature == nil {
		err := errors.New("Invalid feature type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	feature, err := geojson.NewFeaturePb(pb.Feature)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object := &QGeoObject{
		Gid:        pb.Gid,
		Feature:    feature,
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

	if c.Feature == nil {
		err := errors.New("Object feature is not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return c.Feature.Valid()
}

// Pb ...
func (c QGeoObject) Pb() (*proto.QPBxGeoObject, error) {
	pb := &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Properties: c.Properties,
	}

	featurePb, err := c.Feature.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb.Feature = featurePb

	return pb, nil
}
