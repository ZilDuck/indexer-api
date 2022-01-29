package entity

import (
	"fmt"
)

type Nft struct {
	Contract string `json:"contract"`
	TxID     string `json:"txId"`
	BlockNum uint64 `json:"blockNum"`

	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenId  uint64 `json:"tokenId"`
	BaseUri  string `json:"baseUri"`
	TokenUri string `json:"tokenUri"`
	MediaUri string `json:"mediaUri"`

	Zrc1 bool `json:"zrc1"`
	Zrc6 bool `json:"zrc6"`

	Owner string `json:"owner"`

	BurnedAt uint64 `json:"burnedAt,omitempty"`
}

func (n Nft) MetadataUri() string {
	if n.TokenUri != "" {
		return n.TokenUri
	}

	return fmt.Sprintf("%s%d", n.BaseUri, n.TokenId)
}
