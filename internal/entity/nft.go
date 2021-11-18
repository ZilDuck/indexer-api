package entity

type NFT struct {
	Contract string `json:"contract"`
	TokenId  uint64 `json:"tokenId"`
	TokenUri string `json:"tokenUri"`
}
