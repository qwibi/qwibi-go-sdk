package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"google.golang.org/grpc"
	"io"
)

type geoLayer struct {
	layer *layer.QGeoLayer
	*QApiClient
}

func (c *geoLayer) Gid() string {
	if c == nil {
		qlog.Error("Bed layer parameter type nil")
	}

	if c.layer == nil {
		qlog.Error("Bed layer parameter type nil")
	}

	return c.layer.LayerId()
}

func (c *geoLayer) Post(object geo.Object) error {
	req := &proto.QPBxPostRequest{
		LayerId: c.layer.LayerId(),
		Object:  object.Pb(),
	}

	_, err := c.apiClient.Post(c.ctx, req)
	if err != nil {
		return qlog.Error(err.Error())
	}

	return nil
}

// Stream ...
func (c *geoLayer) Stream(handler func(event event.QEvent)) error {
	if c == nil {
		qlog.Error("bed layer parameter type nil")
	}

	if c.layer == nil {
		qlog.Error("Bed layer parameter type nil")
	}

	req := &proto.QPBxStreamRequest{
		LayerId: c.layer.LayerId(),
	}

	client, err := c.apiClient.Stream(c.ctx, req, grpc.EmptyCallOption{})
	if err != nil {
		return qlog.Error(err)
	}

	for {
		in, err := client.Recv()
		if err == io.EOF {
			//qlog.Infof("Client connected: %+v", c)
			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		//qlog.Debugf("Client stream: [%T] %v", in, in)
		event, err := event.NewEventPb(in.Event)
		if err != nil {
			return qlog.Error(err)
		}

		handler(event)
	}
}
