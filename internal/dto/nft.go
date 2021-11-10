package dto

type NFTs map[string]NFT

type NFT map[uint64]Token

type Token struct {
	Uri string `json:"uri"`
}
