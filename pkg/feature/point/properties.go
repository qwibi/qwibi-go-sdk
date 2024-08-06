package point

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QPointProperties struct {
	fid       string
	imageName string
}

func NewPointProperties(options ...PointOption) *QPointProperties {
	h := &QPointProperties{}
	for _, opt := range options {
		opt(h)
	}

	return h
}

func NewPointPropertiesPb(in *proto.QPBxPointFeature_Properties) *QPointProperties {
	return NewPointProperties(
		WithImageName(in.Image),
	)
}

func (c *QPointProperties) Pb() *proto.QPBxPointFeature_Properties {
	return &proto.QPBxPointFeature_Properties{
		Image: c.imageName,
	}
}

type PointOption func(c *QPointProperties)

func WithFid(fid string) PointOption {
	return func(c *QPointProperties) {
		c.fid = fid
	}
}

func WithProperties(properties []byte) PointOption {
	return func(c *QPointProperties) {
		json.Unmarshal(properties, c)
	}
}

func WithImageName(name string) PointOption {
	return func(c *QPointProperties) {
		c.imageName = name
	}
}

//
//func WithPointGid(gid string) PointOption {
//	return func(c *QGeoPoint) {
//		if gid != "" {
//			c.gid = gid
//		}
//	}
//}
//
//func WithPointGeometry(geometry *geometry.QPoint) PointOption {
//	return func(c *QGeoPoint) {
//		c.feature = geometry
//	}
//}
//
//func WithPointProperties(properties []byte) PointOption {
//	return func(c *QGeoPoint) {
//		c.properties = properties
//	}
//}
//
//func WithPointCoordinates(coordinates ...float64) PointOption {
//	return func(c *QGeoPoint) {
//		c.feature.Coordinates = coordinates
//	}
//}
