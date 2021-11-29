package dto

type Contract struct {
	Address  string   `json:"contract"`
	TokenIds []uint64 `json:"tokenIds"`
}
