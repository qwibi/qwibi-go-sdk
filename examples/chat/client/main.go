package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
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

	err = layer.Stream(func(event event.QEvent) {
		qlog.Infof("Event: %+v", event)
		point := geo.NewGeoPoint()
		layer.Post(point)
	})

	if err != nil {
		qlog.Error(err)
	}
}
