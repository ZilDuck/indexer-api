package entity

type ContractState struct {
	Address string                 `json:"address"`
	State   []ContractStateElement `json:"state"`
}