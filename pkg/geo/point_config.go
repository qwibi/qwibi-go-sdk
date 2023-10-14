package geo

type PointOption func(config *QGeoPoint)

func PointGid(gid string) PointOption {
	return func(c *QGeoPoint) {
		c.gid = gid
	}
}

func PointProperties(properties []byte) PointOption {
	return func(c *QGeoPoint) {
		c.properties = properties
	}
}

func PointCoordinates(coordinates ...float64) PointOption {
	return func(c *QGeoPoint) {
		c.geometry.Coordinates = coordinates
	}
}
