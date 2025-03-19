package search

import "github.com/qwibi/qwibi-go-sdk/pkg/qlog"

type SearchOption func(config *QSearch) error

func WithText(text string) SearchOption {
	return func(c *QSearch) error {
		if text == "" {
			return qlog.Error("empty search request")
		} else {
			c.Text = text
		}
		return nil
	}
}
