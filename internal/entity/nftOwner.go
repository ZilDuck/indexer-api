package entity

type NftOwner struct {
	Address  string   `json:"contract"`
	TokenIds []uint64 `json:"tokenIds"`
}
