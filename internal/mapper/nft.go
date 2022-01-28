package mapper

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func NftEntitysToDtos(e []entity.Nft) []dto.Nft {
	nfts := make([]dto.Nft, 0)
	for idx := range e {
		nfts = append(nfts, NftEntityToDto(e[idx]))
	}

	return nfts
}

func NftEntityToDto(e entity.Nft) dto.Nft {
	nft := dto.Nft{
		Contract: e.Contract,
		Name:     e.Name,
		Symbol:   e.Symbol,
		TokenId:  e.TokenId,
		TokenUri: e.TokenUri,
		Owner:    e.Owner,
		BurnedAt: e.BurnedAt,
	}

	if e.Zrc1 == true {
		nft.Type = dto.Zrc1
	}

	if e.Zrc6 == true {
		nft.Type = dto.Zrc6
		nft.TokenUri = fmt.Sprintf("%s%d", e.TokenUri, e.TokenId)
	}

	return nft
}
