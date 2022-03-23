package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func ContractsToDtos(e []entity.Contract) []dto.Contract {
	contracts := make([]dto.Contract, 0)
	for idx := range e {
		contracts = append(contracts, ContractToDto(e[idx]))
	}

	return contracts
}

func ContractToDto(e entity.Contract) dto.Contract {
	contract := dto.Contract{
		Name:            e.Name,
		Address:         e.Address,
		BlockNum:        e.BlockNum,
		Data:            mapData(e.Data),
		MutableParams:   mapParams(e.MutableParams),
		ImmutableParams: mapParams(e.ImmutableParams),
		Transitions:     mapContractTransitions(e.Transitions),
		Standards:       e.Standards,
	}

	return contract
}

func mapData(eData entity.Data) (data dto.Data) {
	data.Tag = eData.Tag
	data.Params = mapParams(eData.Params)
	return
}

func mapParams(eParams []entity.Param) (params []dto.Param) {
	for _, eParam := range eParams {
		param := dto.Param{Type: eParam.Type, VName: eParam.VName}
		if eParam.Value != nil {
			if eParam.Value.Primitive != nil {
				param.Value = eParam.Value.Primitive
			} else {
				param.Value = eParam.Value
			}
		}
		params = append(params, param)
	}
	return
}

func mapContractTransitions(e []entity.ContractTransition) (trans []dto.ContractTransition) {
	for idx := range e {
		t := dto.ContractTransition{
			Name: e[idx].Name,
			Arguments: map[string]string{},
		}
		for _, arg := range e[idx].Arguments {
			t.Arguments[arg.Key] = arg.Value
		}
		trans = append(trans, t)
	}
	return
}
