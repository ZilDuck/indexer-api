package factory

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func NftsIndexToDto(nfts []entity.NFT) dto.NFTs {
	dtos := dto.NFTs{}

	for _, nft := range nfts {
		if _, ok := dtos[nft.Contract]; !ok {
			dtos[nft.Contract] = dto.NFT{}
		}
		dtos[nft.Contract][nft.TokenId] = dto.Token{Uri: nft.TokenUri}
	}

	return dtos
}
