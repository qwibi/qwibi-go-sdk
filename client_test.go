package sdk

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"testing"
)

var _client *QApiClient

func TestNewApiClient(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

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
	log.Info().Msgf("Auth response: %v", res)
}

func TestBasicAuth(t *testing.T) {
	res, err := _client.BasicAuth("user", "password")
	if err != nil {
		t.Fatal(err)
	}
	log.Info().Msgf("Auth response: %v", res)
}

func TestJoin(t *testing.T) {
	layer, err := _client.Join("123")
	if err != nil {
		t.Fatal(err)
	}
	log.Info().Msgf("Layer response: %v", layer)
}
