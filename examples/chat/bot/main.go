package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
)

var client *qwibi.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("Connecto to...", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("Auth with Session... %+v", session)

	layer, err := client.Layer(layer.WithLayerGid("123"))
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("Layer: %+v", layer)

	//
	bot, err := client.Bot(layer.Gid())
	if err != nil {
		qlog.Error(err)
	}

	err = bot.Subscribe(func(request command.QRequest) {
		qlog.Infof("bot request: %+v", request)
		response := command.QResponse{
			Path: "/a/b/c/",
		}
		qlog.Infof("bot response: %+v", response)
		bot.Publish(response)
	})

	if err != nil {
		qlog.Error(err)
	}
}
