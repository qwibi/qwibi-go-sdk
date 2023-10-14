package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"google.golang.org/grpc"
	"io"
	"runtime"
)

func (c *QApiClient) Bot(gid string, handler func(request *command.QRequest)) error {
	client, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
	if err != nil {
		return qlog.Error(err)
	}

	for {
		in, err := client.Recv()
		if err == io.EOF {
			qlog.Infof("Client connected: %+v", c)
			qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		qlog.Infof("==> Bot: [%T] %+v", in, in)

		cmd := &command.QRequest{
			Path: gid,
		}

		handler(cmd)
	}
}
