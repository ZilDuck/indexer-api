package entity

type Contract struct {
	Name            string               `json:"name"`
	Address         string               `json:"address"`
	BlockNum        uint64               `json:"blockNum"`
	Code            string               `json:"code"`
	Data            Data                 `json:"data"`
	MutableParams   Params               `json:"mutableParams"`
	ImmutableParams Params               `json:"immutableParams"`
	Transitions     []ContractTransition `json:"transitions"`
	Standards       map[ZrcStandard]bool `json:"standards"`
}

type ContractTransition struct {
	Index     int                          `json:"index"`
	Name      string                       `json:"name"`
	Arguments []ContractTransitionArgument `json:"arguments"`
}

type ContractTransitionArgument struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ZrcStandard string

const (
	ZRC1 ZrcStandard = "ZRC1"
	ZRC2 ZrcStandard = "ZRC2"
	ZRC3 ZrcStandard = "ZRC3"
	ZRC4 ZrcStandard = "ZRC4"
	ZRC6 ZrcStandard = "ZRC6"
)

type ContractStateElement struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c Contract) MatchesStandard(standard ZrcStandard) bool {
	if _, ok := c.Standards[standard]; ok {
		return true
	}
	return false
}