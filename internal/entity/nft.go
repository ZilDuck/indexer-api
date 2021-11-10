package entity

type NFT struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TxID     string `json:"txId"`
	BlockNum uint64 `json:"blockNum"`

	TokenId  uint64 `json:"tokenId"`
	TokenUri string `json:"tokenUri"`

	By              string `json:"by"`
	ByBech32        string `json:"byBech32"`
	Recipient       string `json:"recipient"`
	RecipientBech32 string `json:"recipientBech32"`
	Owner           string `json:"owner"`
	OwnerBech32     string `json:"ownerBech32"`
}
