package geo

//
//// QGeoPoint ...
//type QGeoPoint struct {
//	gid        string                 `json:"gid"`
//	feature    *feature.QPointFeature `json:"feature"`
//	properties []byte                 `json:"properties"`
//}
//
//// NewGeoPoint ...
//func NewGeoPoint(options ...PointOption) *QGeoPoint {
//	h := &QGeoPoint{
//		gid:     utils.NewID(),
//		feature: feature.NewPointFeature(0, 0),
//	}
//
//	for _, opt := range options {
//		opt(h)
//	}
//
//	return h
//}
//
//func (c *QGeoPoint) Gid() string {
//	return c.gid
//}
//
//func (c *QGeoPoint) Features() *geometry.QGeometry {
//	return geometry.NewGeometry(c.feature)
//}
//
//func (c *QGeoPoint) Properties() []byte {
//	return c.properties
//}
//
//func (c *QGeoPoint) Pb() *proto.QPBxGeoObject {
//	return &proto.QPBxGeoObject{
//		Gid:        c.gid,
//		Geometry:   c.feature.Pb(),
//		Properties: c.properties,
//	}
//}
//
//func (c *QGeoPoint) Valid() error {
//	if c.gid == "" || c.feature == nil || c.feature.Valid() != nil {
//		return qlog.Error("validation error")
//	}
//
//	return nil
//}
