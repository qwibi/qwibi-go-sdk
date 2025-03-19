package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/search"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QSearchRequest struct {
	RequestId string
	Search    *search.QSearch
}

func NewSearchRequest(options ...search.SearchOption) (*QSearchRequest, error) {
	search, err := search.NewSearch(options...)
	if err != nil {
		return nil, qlog.Error(err)
	}
	req := &QSearchRequest{
		RequestId: utils.RequestId(),
		Search:    search,
	}

	return req, nil
}

func NewSearchRequestPb(in *proto.QPBxSearchRequest) (*QSearchRequest, error) {
	return NewSearchRequest(
		search.WithText(in.Search.Text),
	)
}

func (c *QSearchRequest) Pb() *proto.QPBxSearchRequest {
	return &proto.QPBxSearchRequest{
		RequestId: c.RequestId,
		Search:    c.Search.Pb(),
	}
}
