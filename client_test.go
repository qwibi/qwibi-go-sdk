package sdk

import (
	"context"
	"testing"

	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/qlog"
	"github.com/rs/zerolog/log"
)

var client *QApiClient

func init() {
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	addr := "127.0.0.1:8080"
	ctx := context.Background()
	c, err := NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	client = c
}

func TestAnonymousAuth(t *testing.T) {
	res, err := client.AnonymousAuth()
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Anonymous auth response: %#v", res)
}

func TestBasicAuth(t *testing.T) {
	res, err := client.BasicAuth("user", "password")
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Basic auth response: %#v", res)
}

func TestJoinWithoutGid(t *testing.T) {
	layer, err := client.Join("")
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Layer response: %#v", layer)
}

func TestPost(t *testing.T) {

	res1, err := client.Join("")
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Join response: %#v", res1)

	object := geo.NewGeoPoint()
	res2, err := client.Post(object)
	if err != nil {
		t.Fatal(err)
	}

	// qlog.Infof("Post response: %#v", res2)
	qlog.Infof("Post response: %#v", res2)
}

func TestJoinWithGid(t *testing.T) {
	layer, err := client.Join("dXIUiEFyn9WUnh5q")
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Join response: %#v", layer)
}

func TestStream(t *testing.T) {
	log.Info().Msg("Test stream connection")
	err := client.Stream()
	if err != nil {
		t.Fatal(err)
	}
}
