package point

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPointFeature ...
type QPointFeature struct {
	fid        string                   `json:"fid"`
	geometry   *geometry.QPoint         `json:"geometry"`
	properties *QPointFeatureProperties `json:"properties"`
}

// NewPointFeature ...
func NewPointFeature(geometry *geometry.QPoint, option ...PointOption) (*QPointFeature, error) {
	h := &QPointFeature{
		fid:        utils.NewID(),
		geometry:   geometry,
		properties: NewPointProperties(),
	}

	for _, opt := range option {
		opt(h)
	}

	return h, nil
}

// NewPointFeaturePb ...
func NewPointFeaturePb(in *proto.QPBxPointFeature) (*QPointFeature, error) {
	if in == nil {
		err := errors.New("invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	geometry, err := geometry.NewPointPb(in.Geometry)
	if err != nil {
		return nil, qlog.Error(err)
	}

	properties := NewPointFeaturePropertiesPb(in.Properties)

	feature, err := NewPointFeature(geometry,
		WithFid(in.Fid),
		WithProperties(properties),
	)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return feature, feature.Valid()
}

func (c *QPointFeature) Fid() string {
	return c.fid
}

func (c *QPointFeature) Geometry() *geometry.QGeometry {
	return &geometry.QGeometry{c.geometry}
}

func (c *QPointFeature) Properties() []byte {
	b, err := json.Marshal(c.properties)
	if err != nil {
		qlog.Errorf("error marshalling properties: %v", err)
		return nil
	}

	return b
}

func (c *QPointFeature) Pb() *proto.QPBxPointFeature {

	return &proto.QPBxPointFeature{
		Fid:        c.fid,
		Geometry:   c.geometry.Pb(),
		Properties: c.properties.pb(),
	}
}

func (c *QPointFeature) FeaturePb() *proto.QPBxFeature {
	return &proto.QPBxFeature{
		Type: &proto.QPBxFeature_Point{
			Point: c.Pb(),
		},
	}
}

// Valid ...
func (c *QPointFeature) Valid() error {
	if c.geometry == nil {
		return errors.New("invalid parameter geometry")
	}

	return nil
}

//func (c *QPointFeature) MarshalGeoJSON() ([]byte, error) {
//	//type Alias QPointFeature // Create an alias to avoid infinite recursion
//	return json.Marshal(&struct {
//		Type        string    `json:"type"`
//		Coordinates []float64 `json:"coordinates"`
//		//Alias
//	}{
//		Type:        QPointGeometryType,
//		Coordinates: c.Coordinates,
//		//Alias:       (Alias)(*c),
//	})
//}
//
//func (c *QPointFeature) String() string {
//	b, _ := c.MarshalGeoJSON()
//	return string(b)
//}

//var g QGeometry
//g := QPointFeature{}

//func (f QPointFeature) UnmarshalJSON(data []byte) error {
//	qlog.TODO("#### Point")
//	return nil
//}

//func (c *QPointFeature) Value() (driver.Value, error) {
//	return protobuf.Marshal(c.Pb())
//}
//
//func (c *QPointFeature) Scan(src any) error {
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
//func (p *QPointFeature) Scan(value any) error {
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
//func (p QPointFeature) Value() (driver.Value, error) {
//	op := orb.Point{1, 1}
//	return wkb.Value(op).Value()
//}
//
//// FormatParam implements the sql.ParamFormatter interface to tell the SQL
//// builder that the placeholder for a Point parameter needs to be formatted.
//func (p QPointFeature) FormatParam(placeholder string, info *sql.StmtInfo) string {
//	if info.Dialect == dialect.MySQL {
//		return "ST_GeomFromWKB(" + placeholder + ")"
//	}
//	return placeholder
//}
//
//// SchemaType defines the schema-type of the Point object.
//func (QPointFeature) SchemaType() map[string]string {
//	return map[string]string{
//		dialect.MySQL: "POINT",
//	}
//}
