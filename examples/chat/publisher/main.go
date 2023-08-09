package main

import (
	"context"
	"fmt"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/object"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
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
	qlog.Info("Connecto to...", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("Auth with Session... %+v", session)

	//layer, err := client.Stream(gid)
	//if err != nil {
	//	qlog.Error(err)
	//	return
	//}

	//point, err := client.Object(
	//	qwibi.WithGeometry(geometry.NewPoint()),
	//)
	//if err != nil {
	//	return
	//}

	//object, err := client.Object(
	//	qwibi.WithGid("111"),
	//	qwibi.WithGeometry(geometry.NewPoint()),
	//	qwibi.WithProperties([]byte("object properties")),
	//)
	//if err != nil {
	//	qlog.Error(err)
	//}

	object := object.NewGeoObject(
		object.WithGid("myID"),
		object.WithGeometry(geometry.NewPoint()),
		object.WithProperties([]byte("object properties")),
	)

	err = client.Post(object)
	if err != nil {
		qlog.Error(err)
	}

	go func() {
		for {
			object.Properties = []byte(fmt.Sprintf("%s", time.Now()))
			qlog.Infof("Post object: %v", object)

			err = client.Post(object)
			if err != nil {
				qlog.Error(err)
			}

			//client.Connect(gid, object.Gid)

			time.Sleep(3 * time.Second)
		}
	}()

	qlog.Infof("Stream to object: %+v", object)
	err = client.Stream(object.Gid, func(event event.QEvent) {
		qlog.Infof("Event: [%T] %s", event, event)
	})

	if err != nil {
		qlog.Error(err)
	}
}
