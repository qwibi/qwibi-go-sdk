package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
)

type geoBot struct {
	layer *layer.QGeoLayer
	*QApiClient
}

func (c *geoBot) Subscribe(func(request command.QRequest)) error {
	return qlog.TODO()
}

func (c *geoBot) Publish(response command.QResponse) error {
	return qlog.TODO()
}

//// Bot ...
//func (c *geoLayer) Bot(handler func(response *command.QResponse)) error {
//	if c == nil {
//		qlog.Error("bed layer parameter type nil")
//	}
//
//	if c.layer == nil {
//		qlog.Error("Bed layer parameter type nil")
//	}
//
//	//req := &proto.QPBxStreamRequest{
//	//	LayerId: c.layer.LayerId(),
//	//}
//
//	client, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
//	if err != nil {
//		return qlog.Error(err)
//	}
//
//	for {
//		in, err := client.Recv()
//		if err == io.EOF {
//			//qlog.Infof("Client connected: %+v", c)
//			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
//			return qlog.Error(err)
//		}
//		if err != nil {
//			return qlog.Error(err)
//		} b
//
//		//qlog.Debugf("Client stream: [%T] %v", in, in)
//		event, err := event.NewEventPb(in.Event)
//		if err != nil {
//			return qlog.Error(err)
//		}
//
//		handler(event)
//	}
//}
