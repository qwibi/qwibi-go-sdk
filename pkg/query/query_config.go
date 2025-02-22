package query

type QueryOption func(config *QQuery) error

//func WithLayerId(layerId string) QueryOption {
//	return func(c *QQuery) error {
//		if layerId == "" {
//			return qlog.Error("layer ID is not defined")
//		} else {
//			c.PrivateId = layerId
//		}
//		return nil
//	}
//}
