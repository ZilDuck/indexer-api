package dto

type NftOwner struct {
	Address  string   `json:"contract"`
	ZRC6     bool     `json:"zrc6"`
	TokenIds []uint64 `json:"tokenIds,omitempty"`
	NFTs     []Nft    `json:"nfts,omitempty"`
}