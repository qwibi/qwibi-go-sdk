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

	layerResponse, err := client.Layer(
		sdkLayer.WithLayerGid("chat"),
	)
	if err != nil {
		qlog.Error(err)
		return
	}

	qlog.Infof("layer: %+v", layerResponse)

	for {
		command, err := sdkCommand.NewCommand("/message", "hello world")
		if err != nil {
			qlog.Error(err)
			return
		}

		qlog.Infof("command request: %+v", command)

		commandResponse, err := client.Command(layerResponse.Layer.LayerId, command)
		if err != nil {
			qlog.Error(err)
			return
		}

		qlog.Infof("command response: %+v", commandResponse.Response)

		time.Sleep(3 * time.Second)
	}
}
