package search

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QSearch struct {
	Text string `json:"text"`
}

func NewSearch(options ...SearchOption) (*QSearch, error) {
	h := &QSearch{}

	for _, opt := range options {
		err := opt(h)
		if err != nil {
			return nil, qlog.Error(err)
		}
	}

	err := h.Valid()
	if err != nil {
		return nil, qlog.Error()
	}

	return h, nil
}

func NewSearchPb(in *proto.QPBxSearch) (*QSearch, error) {
	return NewSearch(
		WithText(in.Text),
	)
}

func (c *QSearch) Pb() *proto.QPBxSearch {
	return &proto.QPBxSearch{
		Text: c.Text,
	}
}

func (c *QSearch) Valid() error {
	if c.Text == "" {
		return qlog.Error("search request is empty")
	}

	return nil
}
