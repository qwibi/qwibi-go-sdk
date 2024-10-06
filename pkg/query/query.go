package query

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QQuery struct {
	Location string `json:"location"`
}

func NewQuery(options ...QueryOption) (*QQuery, error) {
	h := &QQuery{}

	for _, opt := range options {
		err := opt(h)
		if err != nil {
			return nil, qlog.Error(err)
		}
	}

	return h, nil
}

func NewQueryPb(in *proto.QPBxQuery) (*QQuery, error) {
	return NewQuery()
}
