package point

type PointOption func(c *QPointFeature) error

func WithFid(fid string) PointOption {
	return func(c *QPointFeature) error {
		if fid != "" {
			c.fid = fid
		}
		return nil
	}
}

func WithProperties(properties *QPointFeatureProperties) PointOption {
	return func(c *QPointFeature) error {
		c.properties = properties
		return nil
	}
}
