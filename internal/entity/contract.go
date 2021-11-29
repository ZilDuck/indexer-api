package entity

type Contract struct {
	Name            string   `json:"name"`
	Address         string   `json:"address"`
	AddressBech32   string   `json:"addressBech32"`
	BlockNum        uint64   `json:"blockNum"`
	Code            string   `json:"code"`
	Data            Data     `json:"data"`
	MutableParams   Params   `json:"mutableParams"`
	ImmutableParams Params   `json:"immutableParams"`
	Transitions     []string `json:"transitions"`
	ZRC1            bool     `json:"zrc1"`
	ZRC6            bool     `json:"zrc6"`
}
