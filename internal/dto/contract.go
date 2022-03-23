package dto

import "github.com/ZilDuck/indexer-api/internal/entity"

type Contract struct {
	Name            string                      `json:"name"`
	Address         string                      `json:"address"`
	BlockNum        uint64                      `json:"blockNum"`
	Data            Data                        `json:"data"`
	MutableParams   []Param                     `json:"mutableParams"`
	ImmutableParams []Param                     `json:"immutableParams"`
	Transitions     []ContractTransition        `json:"transitions"`
	Standards       map[entity.ZrcStandard]bool `json:"shape"`
}

type ContractTransition struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments"`
}

type Data struct {
	Tag    string  `json:"_tag,omitempty"`
	Params []Param `json:"params,omitempty"`
}

type Param struct {
	VName string      `json:"vname"`
	Type  string      `json:"type"`
	Value interface{} `json:"value,omitempty"`
}