package geo

import "github.com/qwibi/qwibi-go-sdk/pkg/geometry"

type PointOption func(config *QGeoPoint)

func WithPointGid(gid string) PointOption {
	return func(c *QGeoPoint) {
		c.gid = gid
	}
}

func WithPointGeometry(geometry *geometry.QPoint) PointOption {
	return func(c *QGeoPoint) {
		c.geometry = geometry
	}
}

func WithPointProperties(properties []byte) PointOption {
	return func(c *QGeoPoint) {
		c.properties = properties
	}
}

func WithPointCoordinates(coordinates ...float64) PointOption {
	return func(c *QGeoPoint) {
		c.geometry.Coordinates = coordinates
	}
}
