package zilkroad

type Transaction struct {
	ID           string
	Version      string
	Nonce        string
	Amount       string
	GasPrice     string
	GasLimit     string
	Signature    string
	SenderPubKey string `json:"SenderPubKey"`
	ToAddr       string `json:"ToAddr"`
	Status       State  `json:"Status"`
	Priority     bool   `json:"Priority"`

	Code            string             `json:"Code,omitempty"`
	Data            interface{}        `json:"Data,omitempty"`
	ContractAddress string             `json:"ContractAddress,omitempty"`
	Receipt         TransactionReceipt `json:"Receipt"`
	BlockNum        uint64             `json:"BlockNum"`
}
