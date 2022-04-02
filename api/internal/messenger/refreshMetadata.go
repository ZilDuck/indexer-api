package messenger

type RefreshMetadata struct {
	Contract string `json:"contract"`
	TokenId  uint64 `json:"tokenId"`
	Sent     bool   `json:"state,omitempty"`
}
