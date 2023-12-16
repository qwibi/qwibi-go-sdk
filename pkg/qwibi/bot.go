package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"google.golang.org/grpc"
	"io"
)

type geoBot struct {
	layer *layer.QGeoLayer
	*QApiClient
}

func (c *geoBot) Subscribe(handler func(request *command.QRequest)) error {
	if c == nil {
		qlog.Error("bad bot parameter type nil")
	}

	if c.layer == nil {
		qlog.Error("bad layer parameter type nil")
	}

	bot, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
	if err != nil {
		return qlog.Error(err)
	}

	for {
		in, err := bot.Recv()
		if err == io.EOF {
			//qlog.Infof("Client connected: %+v", c)
			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		//qlog.Debugf("Client stream: [%T] %v", in, in)
		command, err := command.NewRequestPb(in.Command)
		if err != nil {
			return qlog.Error(err)
		}

		handler(command)
	}
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
