package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"strconv"
)

var supportedActions = []string{"mint", "transfer", "burn", "sale", "listing", "delisting"}

func ActionsToDtos(e []entity.NftAction) []dto.NftAction {
	actions := make([]dto.NftAction, 0)
	for idx := range e {
		action := ActionToDto(e[idx])
		if action != nil {
			actions = append(actions, *action)
		}
	}

	return actions
}

func ActionToDto(e entity.NftAction) *dto.NftAction {
	matched := false
	for _, supportedAction := range supportedActions {
		if e.Action == supportedAction {
			matched = true
		}
	}
	if matched == false {
		return nil
	}

	a := &dto.NftAction{
		BlockNum: e.BlockNum,
		TxID: e.TxID,
		Action: e.Action,
	}

	if e.From != "" {
		a.From = &e.From
	}

	if e.To != "" {
		a.To = &e.To
	}

	if e.Marketplace != "" {
		a.Marketplace = &e.Marketplace
	}

	if e.Cost != "" {
		n, err := strconv.ParseUint(e.Cost, 10, 64)
		if err == nil {
			a.Cost = &n
		}
	}

	if e.Fee != "" {
		n, err := strconv.ParseUint(e.Fee, 10, 64)
		if err == nil {
			a.Fee = &n
		}
	}

	if e.Royalty != "" {
		n, err := strconv.ParseUint(e.Royalty, 10, 64)
		if err == nil {
			a.Royalty = &n
		}
	}

	if e.Fungible != "" {
		a.Fungible = &e.Fungible
	}

	return a
}
