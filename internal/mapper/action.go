package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func ActionsToDtos(e []entity.NftAction) []dto.NftAction {
	nfts := make([]dto.NftAction, 0)
	for idx := range e {
		nfts = append(nfts, ActionToDto(e[idx]))
	}

	return nfts
}

func ActionToDto(e entity.NftAction) dto.NftAction {
	return dto.NftAction{
		BlockNum: e.BlockNum,
		TxID: e.TxID,
		Action: e.Action,
		From: e.From,
		To: e.To,
	}
}
