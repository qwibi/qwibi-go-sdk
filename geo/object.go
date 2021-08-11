package geo

import (
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geojson"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//QGeoObject ...
type QGeoObject struct {
	Gid        string
	OwnerID    string
	LayerID    string
	Geometry   geojson.QGeometry
	Properties *structpb.Struct
	Ctime      *timestamppb.Timestamp
}

func (c *QGeoObject) Valid() error {
	if c.Gid == "" || c.OwnerID == "" || c.LayerID == "" {
		return errors.New("Invalid format")
	}

	return nil
}

// NewGeoObjectPb ...
func NewGeoObjectPb(pb *proto.QPBxGeoObject) (*QGeoObject, error) {

	if pb == nil {
		err := errors.New("Invalid format type nil")
		logrus.Error(err)
		return nil, err
	}

	geom, err := geojson.NewGeometryPb(pb.Geometry)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	object := &QGeoObject{
		Gid:        pb.Gid,
		OwnerID:    pb.OwnerID,
		LayerID:    pb.LayerID,
		Geometry:   geom,
		Properties: pb.Properties,
		Ctime:      pb.Ctime,
	}

	if err := object.Valid(); err != nil {
		return nil, err
	}

	return object, nil
}

// NewGeoObject ...
func NewGeoObject(ownerID string, layerID string, feature *geojson.QFeature) (*QGeoObject, error) {

	if feature == nil {
		err := errors.New("Invalid format type nil")
		logrus.Error(err)
		return nil, err
	}

	object := &QGeoObject{
		Gid:        feature.Gid,
		OwnerID:    ownerID,
		LayerID:    layerID,
		Geometry:   feature.Geometry,
		Properties: feature.Properties,
		Ctime:      timestamppb.Now(),
	}

	return object, nil
}

// NewObjectPb ...
func NewObjectPb(pb *proto.QPBxGeoObject) (*QGeoObject, error) {

	if pb == nil {
		err := errors.New("Invalid format type nil")
		logrus.Error(err)
		return nil, err
	}

	pbGeometry, err := geojson.NewGeometryPb(pb.Geometry)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	object := &QGeoObject{
		Gid:        pb.Gid,
		OwnerID:    pb.OwnerID,
		LayerID:    pb.LayerID,
		Geometry:   pbGeometry,
		Properties: pb.Properties,
		Ctime:      pb.Ctime,
	}

	return object, nil
}

// Pb ...
func (c *QGeoObject) Pb() (*proto.QPBxGeoObject, error) {

	if err := c.Valid(); err != nil {
		return nil, err
	}

	pbGeometry, err := c.Geometry.GeometryPb()
	if err != nil {
		return nil, err
	}

	pb := &proto.QPBxGeoObject{
		Gid:        c.Gid,
		OwnerID:    c.OwnerID,
		LayerID:    c.LayerID,
		Geometry:   pbGeometry,
		Properties: c.Properties,
		Ctime:      c.Ctime,
	}

	return pb, nil
}
