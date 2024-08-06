package query

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QQuery struct {
	LayerId string `json:"layerId"`
}

func NewQuery(layerId string, options ...QueryOption) (*QQuery, error) {
	h := &QQuery{
		LayerId: layerId,
	}

	for _, opt := range options {
		err := opt(h)
		if err != nil {
			return nil, qlog.Error(err)
		}
	}

	return h, nil
}

func NewQueryPb(in *proto.QPBxQuery) (*QQuery, error) {
	if in == nil {
		return nil, qlog.Error("bed parameter type nil")
	}
	return NewQuery(in.LayerId)
}
