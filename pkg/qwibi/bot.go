package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	sdkRequest "github.com/qwibi/qwibi-go-sdk/pkg/rpc/request"
	sdkResponse "github.com/qwibi/qwibi-go-sdk/pkg/rpc/response"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"io"
)

type geoBot struct {
	client proto.QPBxApi_BotClient
}

func (c *geoBot) Subscribe(handler func(request *sdkRequest.QCommandRequest)) error {
	if c.client == nil {
		return qlog.Error("bad client parameter type nil")
	}

	for {
		in, err := c.client.Recv()
		if err == io.EOF {
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		request, err := sdkRequest.NewCommandRequestPb(in)
		if err != nil {
			return qlog.Error(err)
		}

		handler(request)
	}
}

func (c *geoBot) Publish(requestId string, layerId string, response *command.QResponse) error {
	if c.client == nil || response == nil {
		return qlog.Error("bad client parameter type nil")
	}

	res, err := sdkResponse.NewCommandResponse(requestId, layerId, response)
	if err != nil {
		return qlog.Error(err)
	}

	err = c.client.Send(res.Pb())
	if err != nil {
		return qlog.Error(err)
	}

	return nil
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
