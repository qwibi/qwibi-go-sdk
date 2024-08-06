package geometry

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPoint ...
type QPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

// NewPoint ...
func NewPoint(coordinates ...float64) (*QPoint, error) {
	point := &QPoint{
		Coordinates: coordinates,
	}

	return point, nil
}

// NewPointPb ...
func NewPointPb(in *proto.QPBxPoint) (*QPoint, error) {
	if in == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	point, err := NewPoint(in.Coordinates...)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return point, nil
}

//func NewPointMap(f map[string]interface{}) (QPoint, error) {
//	for _, i := range f {
//		switch v := i.(type) {
//		default:
//			qlog.Errorf("@@@@@@@@ %s", v)
//			qlog.Errorf("@@@@ %T", i)
//		}
//	}
//	return nil, qlog.TODO()
//}

// GetType...
func (c *QPoint) GetType() string {
	return QPointGeometryType
}

// Valid ...
func (c *QPoint) Valid() error {
	if len(c.Coordinates) < 2 {
		err := errors.New("Wrong coordinates format")
		return qlog.Error(err)
	}

	return nil
}

// Pb ...
func (c *QPoint) Pb() *proto.QPBxPoint {
	return &proto.QPBxPoint{
		Coordinates: c.Coordinates,
	}
}

func (c *QPoint) MarshalGeoJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}{
		Type:        QPointGeometryType,
		Coordinates: c.Coordinates,
	})
}

//func (c *QPoint) String() string {
//	b, _ := c.MarshalGeoJSON()
//	return string(b)
//}

//var g QGeometry
//g := QPoint{}

//func (f QPoint) UnmarshalJSON(data []byte) error {
//	qlog.TODO("#### Point")
//	return nil
//}

//func (c *QPoint) Value() (driver.Value, error) {
//	return protobuf.Marshal(c.Pb())
//}
//
//func (c *QPoint) Scan(src any) error {
//	if src == nil {
//		return nil
//	}
//
//	pb := &proto.QPBxPoint{}
//
//	if b, ok := src.([]byte); ok {
//		if err := protobuf.Unmarshal(b, pb); err != nil {
//			return err
//		}
//		return nil
//	}
//
//	c.Coordinates = pb.Coordinates
//
//	return fmt.Errorf("unexpected type %T", src)
//}

//// Scan implements the Scanner interface.
//func (p *QPoint) Scan(value any) error {
//	bin, ok := value.([]byte)
//	if !ok {
//		return fmt.Errorf("invalid binary value for point")
//	}
//	var op orb.Point
//	if err := wkb.Scanner(&op).Scan(bin[4:]); err != nil {
//		return err
//	}
//	//p[0], p[1] = op.X(), op.Y()
//	return nil
//}
//
//// Value implements the driver Valuer interface.
//func (p QPoint) Value() (driver.Value, error) {
//	op := orb.Point{1, 1}
//	return wkb.Value(op).Value()
//}
//
//// FormatParam implements the sql.ParamFormatter interface to tell the SQL
//// builder that the placeholder for a Point parameter needs to be formatted.
//func (p QPoint) FormatParam(placeholder string, info *sql.StmtInfo) string {
//	if info.Dialect == dialect.MySQL {
//		return "ST_GeomFromWKB(" + placeholder + ")"
//	}
//	return placeholder
//}
//
//// SchemaType defines the schema-type of the Point object.
//func (QPoint) SchemaType() map[string]string {
//	return map[string]string{
//		dialect.MySQL: "POINT",
//	}
//}
