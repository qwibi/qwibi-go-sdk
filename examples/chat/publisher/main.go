package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
	"math/rand"
	"time"
)

var client *qwibi.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("connecto to...", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("auth with session... %+v", session)

	layer, err := client.Layer(
	//layer.WithLayerGid("12458"),
	//layer.WithLayerPublic(true),
	)

	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("layer: %+v", layer)

	point := geo.NewGeoPoint(
		geo.WithPointGid("me"),
		geo.WithPointCoordinates(1.1+rand.Float64(), 2.2+rand.Float64()),
	)

	go func() {
		for {
			qlog.Infof("post: %+v", point)
			err = layer.Post(point)
			if err != nil {
				qlog.Error(err)
			}

			time.Sleep(3 * time.Second)
		}
	}()

	err = layer.Stream(func(event event.QEvent) {
		qlog.Infof("event: [%T] %+v", event, event)
	})

	if err != nil {
		qlog.Error(err)
	}
}
