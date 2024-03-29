package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

func NftOwnerToDtos(e []entity.NftOwner) []dto.NftOwner {
	nftOwners := make([]dto.NftOwner, 0)
	for idx := range e {
		nftOwners = append(nftOwners, NftOwnerToDto(e[idx]))
	}

	return nftOwners
}

func NftOwnerToDto(e entity.NftOwner) dto.NftOwner {
	return dto.NftOwner{
		Address: util.ToCheckSumAddress(e.Address),
		ZRC6: e.ZRC6,
		TokenIds: e.TokenIds,
		NFTs: NftToDtos(e.NFTs),
	}
}
