package geo

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeoPoint ...
type QGeoPoint struct {
	gid        string           `json:"gid"`
	geometry   *geometry.QPoint `json:"geometry"`
	properties []byte           `json:"properties"`
}

// NewGeoPoint ...
func NewGeoPoint(options ...PointOption) *QGeoPoint {
	h := &QGeoPoint{
		gid:      utils.NewID(),
		geometry: geometry.NewPoint(),
	}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func (c *QGeoPoint) Gid() string {
	return c.gid
}

func (c *QGeoPoint) Geometry() *geometry.QGeometry {
	return geometry.NewGeometry(c.geometry)
}

func (c *QGeoPoint) Properties() []byte {
	return c.properties
}

func (c *QGeoPoint) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Gid:        c.gid,
		Geometry:   c.geometry.Pb(),
		Properties: c.properties,
	}
}

func (c *QGeoPoint) Valid() error {
	if c.gid == "" || c.geometry == nil || c.geometry.Valid() != nil {
		return qlog.Error("validation error")
	}

	return nil
}
