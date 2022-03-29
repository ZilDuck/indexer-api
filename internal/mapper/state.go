package mapper

import (
	"encoding/json"
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func StateToDto(e entity.ContractState, filters []string) dto.State {
	state := dto.State{}
	for _, el := range e.State {
		if len(filters) > 1 {
			matched := false
			for _, filter :=  range filters {
				if el.Key == filter {
					matched = true
				}
			}
			if !matched {
				continue
			}
		}

		var jsonEl interface{}
		if err := json.Unmarshal([]byte(el.Value), &jsonEl); err == nil {
			state[el.Key] = jsonEl
		} else {
			state[el.Key] = el.Value
		}
	}

	return state
}
