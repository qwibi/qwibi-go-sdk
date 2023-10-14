package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

//type Channel struct {
//	ID   string `json:"id"`
//	Type string `json:"type"`
//	CID  string `json:"cid"` // full id in format channel_type:channel_ID
//	Team string `json:"team"`
//
//	Config ChannelConfig `json:"config"`
//
//	CreatedBy *User `json:"created_by"`
//	Disabled  bool  `json:"disabled"`
//	Frozen    bool  `json:"frozen"`
//
//	MemberCount int              `json:"member_count"`
//	Members     []*ChannelMember `json:"members"`
//
//	Messages        []*Message     `json:"messages"`
//	PinnedMessages  []*Message     `json:"pinned_messages"`
//	PendingMessages []*Message     `json:"pending_messages"`
//	Read            []*ChannelRead `json:"read"`
//
//	CreatedAt     time.Time `json:"created_at"`
//	UpdatedAt     time.Time `json:"updated_at"`
//	LastMessageAt time.Time `json:"last_message_at"`
//
//	TruncatedBy *User      `json:"truncated_by"`
//	TruncatedAt *time.Time `json:"truncated_at"`
//
//	ExtraData map[string]interface{} `json:"-"`
//
//	client *Client
//}

// QGeoLayer ...
type QGeoLayer struct {
	layerId    string `json:"layerId"`
	public     bool   `json:"public"`
	properties []byte `json:"properties"`
}

// NewGeoLayer ...
func NewGeoLayer(options ...LayerOption) *QGeoLayer {
	h := &QGeoLayer{}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func NewGeoLayerPb(in *proto.QPBxGeoLayer) (*QGeoLayer, error) {
	layer := NewGeoLayer(
		WithLayerGid(in.LayerId),
		WithLayerProperties(in.Properties),
	)
	return layer, nil

}

func (c *QGeoLayer) LayerId() string {
	return c.layerId
}

func (c *QGeoLayer) Public() bool {
	return c.public
}

func (c *QGeoLayer) Properties() []byte {
	return c.properties
}

func (c *QGeoLayer) Pb() *proto.QPBxGeoLayer {
	return &proto.QPBxGeoLayer{
		LayerId:    c.layerId,
		Properties: c.properties,
	}
}

func (c *QGeoLayer) Valid() error {
	if c.layerId == "" {
		return qlog.Error("layer validation error")
	}

	return nil
}

//type QGeoLayer struct {
//	layerId string `json:"layerId"`
//	//geometry   *geometry.QLayer `json:"geometry"`
//	//properties []byte `json:"properties"`
//	//client *qwibi.QApiClient
//}
//
////func (c *qwibi.QApiClient) Layer() (*QGeoLayer, error) {
////	graph := &QGeoLayer{
////		client: c,
////	}
////	return graph, nil
////}
//
//// NewGeoLayer ...
//func NewGeoLayer(options ...Option) *QGeoLayer {
//	h := &QGeoLayer{
//		layerId: utils.NewID(),
//	}
//
//	for _, opt := range options {
//		opt(h)
//	}
//
//	return h
//}
//
//func NewLayerPb(in *proto.QPBxLayerRequest) (*QGeoLayer, error) {
//	if in == nil {
//		return nil, qlog.Error("Bad layer request type nil")
//	}
//
//	layer := &QGeoLayer{
//		layerId: in.LayerId,
//	}
//
//	return layer, nil
//}
//
//func (c *QGeoLayer) Pb() *proto.QPBxGeoLayer {
//	return &proto.QPBxGeoLayer{
//		AccountId: c.layerId,
//	}
//}
