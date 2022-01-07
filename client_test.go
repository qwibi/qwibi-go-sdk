package sdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/geojson"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var _client *QApiClient

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func TraceError(err error) {
	out, ok := errors.Cause(err).(stackTracer)
	if !ok {
		fmt.Printf("%v", err)
		panic("Error does not implement stackTracer")
	}
	st := out.StackTrace()
	fmt.Printf("%+v\n", st[0:2])
}

func TestNewApiClient(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// log.Logger = log.With().Caller().Logger()

	addr := "127.0.0.1:8080"

	client, err := NewClient(addr)
	if err != nil {
		t.Fatal(err)
	}

	_client = client
}

func TestAnonymousAuth(t *testing.T) {
	res, err := _client.AnonymousAuth()
	if err != nil {
		t.Fatal(err)
	}
	log.Info().Msgf("Auth response: %#v", res)
}

func TestBasicAuth(t *testing.T) {
	res, err := _client.BasicAuth("user", "password")
	if err != nil {
		t.Fatal(err)
	}
	log.Info().Msgf("Auth response: %#v", res)
}

func TestJoin(t *testing.T) {
	layer, err := _client.Join("123")
	if err != nil {
		t.Fatal(err)
	}
	log.Info().Msgf("Layer response: %#v", layer)
}

func TestPost(t *testing.T) {
	object := &geo.QGeoObject{
		Gid:     "123",
		OwnerID: "123",
		LayerID: "123",
		Feature: geojson.NewPointFeautre(),
	}

	res, err := _client.Post(object)
	if err != nil {
		TraceError(err)
		t.Fatal(err)
	}
	log.Info().Msgf("Post response: %#v", res)
}
