package factory

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func NftsIndexToDto(nfts []entity.NFT) []*dto.Contract {
	contracts := make([]*dto.Contract, 0)

	for _, nft := range nfts {

		var contract *dto.Contract
		for idx := range contracts {
			if contracts[idx].Address == nft.Contract {
				contract = contracts[idx]
			}
		}
		if contract == nil {
			contract = &dto.Contract{Address: nft.Contract}
			contracts = append(contracts, contract)
		}

		contract.TokenIds = append(contract.TokenIds, nft.TokenId)
	}

	return contracts
}
