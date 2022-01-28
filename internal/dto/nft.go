package dto

type NftType string

const (
	Zrc1 NftType = "Zrc1"
	Zrc6 NftType = "Zrc6"
)

type Nft struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`

	TokenId  uint64 `json:"tokenId"`
	TokenUri string `json:"tokenUri"`

	Type NftType `json:"type"`

	Owner string `json:"owner"`

	BurnedAt uint64 `json:"burnedAt,omitempty"`
}
