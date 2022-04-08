package entity

type NftAction struct {
	Contract    string `json:"contract"`
	TokenId     uint64 `json:"tokenId"`
	TxID        string `json:"txId"`
	BlockNum    uint64 `json:"blockNum"`
	Action      string `json:"action"`
	From        string `json:"from"`
	To          string `json:"to"`
	Zrc1        bool   `json:"zrc1"`
	Zrc6        bool   `json:"zrc6"`
	Marketplace string `json:"marketplace"`
	Cost        string `json:"cost"`
	Fee         string `json:"fee"`
	Royalty     string `json:"royalty"`
	Fungible    string `json:"fungible"`
}