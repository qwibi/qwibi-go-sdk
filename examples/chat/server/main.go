package main

import (
	"context"
	sdk "github.com/qwibi/qwibi-go-sdk"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
)

var client *sdk.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = sdk.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}

	qlog.Info("Connecto to...", addr)
	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("Auth with Session... %+v", session)

	layer := geo.NewGeoLayer(
		geo.WithGid("chat"),
	)
	object, err := client.Post(layer)
	qlog.Infof("Post new Layer... %+v", layer)

	qlog.Info("Start Bot...", object.GetGid())

	client.Bot(object.GetGid(), func(request *command.QRequest) {
		qlog.Infof("Bot: request => %s", request)

		//_, err := client.Command("/dodo")
		//if err != nil {
		//	qlog.Error(err)
		//}
	})
}