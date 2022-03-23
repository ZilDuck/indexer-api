package dto

type NftAction struct {
	TxID     string `json:"txId"`
	BlockNum uint64 `json:"blockNum"`
	Action   string `json:"action"`
	From     string `json:"from"`
	To       string `json:"to"`
}