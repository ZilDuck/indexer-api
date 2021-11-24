package entity

type NFT struct {
	Contract string `json:"contract"`
	TxID     string `json:"txId"`
	BlockNum uint64 `json:"blockNum"`

	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenId  uint64 `json:"tokenId"`
	TokenUri string `json:"tokenUri"`

	Zrc1 bool `json:"zrc1"`
	Zrc6 bool `json:"zrc6"`

	Owner string `json:"owner"`

	BurnedAt uint64 `json:"burnedAt"`
}
