package factory

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func NftsIndexToDto(nfts []entity.NFT) dto.NFTs {
	dtos := dto.NFTs{}

	for _, nft := range nfts {
		if _, ok := dtos[nft.Contract]; !ok {
			dtos[nft.Contract] = dto.NFT{}
		}

		var token dto.Token
		switch {
		case nft.Zrc1:
			token.Uri = nft.TokenUri
			token.Type = "ZRC1"
			break
		case nft.Zrc6:
			token.Uri = fmt.Sprintf("%s%d", nft.TokenUri, nft.TokenId)
			token.Type = "ZRC6"
			break
		}

		dtos[nft.Contract][nft.TokenId] = token
	}

	return dtos
}
