package dto

type NftType string

const (
	Zrc1 NftType = "ZRC1"
	Zrc6 NftType = "ZRC6"
)

type Nft struct {
	Contract string      `json:"contract"`
	Name     string      `json:"name"`
	Symbol   string      `json:"symbol"`
	TokenId  uint64      `json:"tokenId"`
	TokenUri string      `json:"tokenUri"`
	Owner    string      `json:"owner"`
	Type     NftType     `json:"type"`
	Metadata interface{} `json:"metadata,omitempty"`
	AssetUri string      `json:"assetUri,omitempty"`
	BurnedAt uint64      `json:"burnedAt,omitempty"`
}
