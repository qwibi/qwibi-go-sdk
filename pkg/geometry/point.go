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
func NewPoint() QPoint {
	return QPoint{
		Coordinates: []float64{0, 0},
	}
}

// NewPointPb ...
func NewPointPb(pb *proto.QPBxPoint) (QPoint, error) {
	point := NewPoint()

	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return point, qlog.Error(err)
	}

	point.Coordinates = pb.Coordinates

	return point, point.Valid()
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

// Valid ...
func (c QPoint) GetType() string {
	return QPointGeometryType
}

// Valid ...
func (c QPoint) Valid() error {
	if len(c.Coordinates) < 2 {
		err := errors.New("Wrong coordinates format")
		return qlog.Error(err)
	}

	return nil
}

// Pb ...
func (c QPoint) Pb() *proto.QPBxGeometry {
	return &proto.QPBxGeometry{
		Type: &proto.QPBxGeometry_Point{
			Point: &proto.QPBxPoint{
				Coordinates: c.Coordinates,
			},
		},
	}
}

func (c QPoint) MarshalJSON() ([]byte, error) {
	//type Alias QPoint // Create an alias to avoid infinite recursion
	return json.Marshal(&struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
		//Alias
	}{
		Type:        QPointGeometryType,
		Coordinates: c.Coordinates,
		//Alias:       (Alias)(*c),
	})
}

//func (f QPoint) UnmarshalJSON(data []byte) error {
//	qlog.TODO("#### Point")
//	return nil
//}
