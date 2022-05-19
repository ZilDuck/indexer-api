package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

func NftToDtos(e []entity.Nft) []dto.Nft {
	nfts := make([]dto.Nft, 0)
	for idx := range e {
		nfts = append(nfts, NftToDto(e[idx]))
	}

	return nfts
}

func NftToDto(e entity.Nft) dto.Nft {
	nft := dto.Nft{
		Contract: util.ToCheckSumAddress(e.Contract),
		Name:     e.Name,
		Symbol:   e.Symbol,
		TokenId:  e.TokenId,
		TokenUri: e.TokenUri,
		Owner:    util.ToCheckSumAddress(e.Owner),
		BurnedAt: e.BurnedAt,
	}
	if e.Metadata != nil && len(e.Metadata.Properties) != 0 {
		nft.Metadata = e.Metadata.Properties
	}

	if e.Zrc1 == true {
		nft.Type = dto.Zrc1
	}

	if e.Zrc6 == true {
		nft.Type = dto.Zrc6
		nft.TokenUri = e.MetadataUri()
	}

	return nft
}
