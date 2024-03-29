package dto

import "github.com/ZilDuck/indexer-api/internal/entity"

type NftAction struct {
	TxID     string            `json:"txId"`
	BlockNum uint64            `json:"blockNum"`
	Action   entity.ActionType `json:"action"`

	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`

	Marketplace *string `json:"marketplace,omitempty"`
	Cost        *uint64 `json:"cost,omitempty"`
	Fee         *uint64 `json:"fee,omitempty"`
	Royalty     *uint64 `json:"royalty,omitempty"`
	Fungible    *string `json:"fungible,omitempty"`
}