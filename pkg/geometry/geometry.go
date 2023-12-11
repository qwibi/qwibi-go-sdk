package geometry

import (
	"database/sql/driver"
	"encoding/hex"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	geom2 "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

// QGeometry ...
type Geometry interface {
	Valid() error
	Pb() *proto.QPBxGeometry
	GetType() string
	//String() string
	//Value() (driver.Value, error)
	//UnmarshalJSON(data []byte) error
}

type QGeometry struct {
	Geometry
}

func NewGeometry(g Geometry) *QGeometry {
	return &QGeometry{g}
}

// NewGeometryPb ...
func NewGeometryPb(in *proto.QPBxGeometry) (*QGeometry, error) {
	if in == nil {
		return nil, qlog.Error("Wrong geometry parameter type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxGeometry_Point:
		point, err := NewPointPb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}
		return &QGeometry{point}, nil
	default:
		return nil, qlog.Errorf("Unknown geometry type: %s", v)
	}
}

func NewGeometryStruct(v map[string]interface{}) (*QGeometry, error) {
	if v == nil {
		return nil, qlog.Error("Bad parameter type nil")
	}

	switch t := v["type"]; t {
	case QPointGeometryType:
		var o QPoint
		err := mapstructure.Decode(v, &o)
		if err != nil {
			return nil, qlog.Error(err)
		}
		return &QGeometry{&o}, nil
	default:
		return nil, qlog.Errorf("Unknown gemetry type: %s", t)
	}

}

func (c *QGeometry) Value() (driver.Value, error) {
	v := c.String()
	return driver.Value(v), nil
}

func (c *QGeometry) Scan(src any) error {
	if src == nil {
		return qlog.Error("scan src is not defined")
	}

	s := fmt.Sprintf("%s", src)

	qlog.Debugf("Scan: %s", s)

	wkbData, err := hex.DecodeString(s)
	if err != nil {
		return qlog.Errorf("error decoding WKB: %v", err)
	}

	// Decode EWKB data
	geom, err := ewkb.Unmarshal(wkbData)
	if err != nil {
		return qlog.Errorf("error unmarshalling WKB: %v", err)
	}

	switch v := geom.(type) {
	case *geom2.Point:
		c.Geometry = NewPoint(v.FlatCoords()...)
		return nil
	default:
		return qlog.Error("unknown geometry type: %T", v)
	}
}

func (c *QGeometry) String() string {
	return fmt.Sprint(c.Geometry)
}

func (c *QGeometry) FormatParam(placeholder string, info *sql.StmtInfo) string {
	if info.Dialect == dialect.Postgres {
		return "ST_GeomFromGeoJSON(" + placeholder + ")"
	}

	return placeholder
}

func (c *QGeometry) SchemaType() map[string]string {
	return map[string]string{
		dialect.Postgres: "geometry",
	}
}
