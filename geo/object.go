package geo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geojson"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// QGeoObject ...
//QGeoObject ...
type QGeoObject struct {
	Gid        string
	OwnerID    string
	LayerID    string
	Feature    *geojson.QFeature
	Properties *structpb.Struct
	Ctime      *timestamppb.Timestamp
}

func (c *QGeoObject) Valid() error {
	if c.Gid == "" || c.OwnerID == "" || c.LayerID == "" || c.Feature == nil {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if err := c.Feature.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// NewGeoObjectPb ...
func NewGeoObjectPb(pb *proto.QPBxGeoObject) (*QGeoObject, error) {

	if pb == nil || pb.Feature == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	feature, err := geojson.NewFeaturePb(pb.Feature)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	object := &QGeoObject{
		Gid:        pb.Gid,
		OwnerID:    pb.OwnerID,
		LayerID:    pb.LayerID,
		Feature:    feature,
		Properties: pb.Properties,
		Ctime:      pb.Ctime,
	}

	if err := object.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return object, nil
}

// NewGeoObject ...
func NewGeoObject(gid string, ownerID string, layerID string, feature *geojson.QFeature) (*QGeoObject, error) {

	if feature == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object := &QGeoObject{
		Gid:        gid,
		OwnerID:    ownerID,
		LayerID:    layerID,
		Feature:    feature,
		Properties: feature.Properties,
		Ctime:      timestamppb.Now(),
	}

	return object, nil
}

// Pb ...
func (c *QGeoObject) Pb() (*proto.QPBxGeoObject, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	featurePb, err := c.Feature.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fmt.Println(">>>>>", c.OwnerID)
	pb := &proto.QPBxGeoObject{
		Gid:     c.Gid,
		OwnerID: c.OwnerID,
		// LayerID:    c.LayerID,
		Feature:    featurePb,
		Properties: c.Properties,
		Ctime:      c.Ctime,
	}

	return pb, nil
}
