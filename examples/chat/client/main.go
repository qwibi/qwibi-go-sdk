package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	sdkCommand "github.com/qwibi/qwibi-go-sdk/pkg/command"
	sdkLayer "github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("connect:", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("auth session: %+v", session)

	layer, err := client.Layer(
		sdkLayer.WithLayerGid("chat"),
	)
	if err != nil {
		qlog.Error(err)
		return
	}

	qlog.Infof("layer: %+v", layer)

	for {
		command, err := sdkCommand.NewCommand("/help")
		if err != nil {
			qlog.Error(err)
			return
		}

		qlog.Infof("command request: %+v", command)

		response, err := client.Command(layer.LayerId, command)
		if err != nil {
			qlog.Error(err)
			return
		}

		qlog.Infof("command response: %+v", response.Response)

		time.Sleep(3 * time.Second)
	}
}
