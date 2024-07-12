package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QStream struct {
	layer *layer.QLayer
	*QApiClient
}

func (c *QStream) Post(object geo.Object) error {
	req := &proto.QPBxPostRequest{
		LayerId: c.layer.LayerId,
		Object:  object.Pb(),
	}

	_, err := c.apiClient.Post(c.ctx, req)
	if err != nil {
		return qlog.Error(err.Error())
	}

	return nil
}
