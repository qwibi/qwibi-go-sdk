package request

import (
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/geojson"
)

// QPostRequest ...
type QPostRequest struct {
	feature geojson.QFeature
}
type QPostResponse struct {
	object geo.QGeoObject
}
