package entity

import (
	"errors"
	"fmt"
)

type Transaction struct {
	ID           string
	Version      string
	Nonce        string
	Amount       string
	GasPrice     string
	GasLimit     string
	Signature    string
	SenderPubKey string
	ToAddr       string
	Status       int
	Priority     bool
	BlockNum     uint64 `json:"BlockNum"`

	Code    string             `json:"Code,omitempty"`
	Data    Data               `json:"Data,omitempty"`
	Receipt TransactionReceipt `json:"Receipt"`

	IsContractCreation    bool   `json:"ContractCreation"`
	IsContractExecution   bool   `json:"ContractExecution"`
	ContractAddress       string `json:"ContractAddress,omitempty"`
	ContractAddressBech32 string `json:"ContractAddressBech32,omitempty"`

	SenderAddr   string `json:"SenderAddr"`
	SenderBech32 string `json:"SenderBech32"`
}

type Data struct {
	Tag    string `json:"_tag,omitempty"`
	Params Params `json:"params,omitempty"`
}

type TransactionReceipt struct {
	Accept        bool                   `json:"accept"`
	Errors        interface{}            `json:"errors"`
	Exceptions    []TransactionException `json:"exceptions"`
	Success       bool                   `json:"success"`
	CumulativeGas string                 `json:"cumulative_gas"`
	EpochNum      string                 `json:"epoch_num"`
	Transitions   []Transition           `json:"transitions"`
	EventLogs     []EventLog             `json:"event_logs"`
}

type TransactionException struct {
	Line    int    `json:"line"`
	Message string `json:"message"`
}

type EventLog struct {
	EventName string `json:"_eventname"`
	Address   string `json:"address"`
	Params    Params `json:"params"`
}

type Transition struct {
	Accept bool              `json:"accept"`
	Addr   string            `json:"addr"`
	Depth  int               `json:"depth"`
	Msg    TransitionMessage `json:"msg"`
}

type TransitionMessage struct {
	Amount  string `json:"_amount"`
	Receipt string `json:"_receipt"`
	Tag     string `json:"_tag"`
	Params  Params `json:"params"`
}

func (tx Transaction) GetEventLogs(eventName string) []EventLog {
	eventLogs := make([]EventLog, 0)
	for _, event := range tx.Receipt.EventLogs {
		if event.EventName == eventName {
			eventLogs = append(eventLogs, event)
		}
	}
	return eventLogs
}

func (tx Transaction) GetEventLogForAddr(addr, eventName string) (EventLog, error) {
	for _, event := range tx.Receipt.EventLogs {
		if event.Address == addr && event.EventName == eventName {
			return event, nil
		}
	}
	return EventLog{}, errors.New(fmt.Sprintf("Event %s for address %s does not exist", eventName, addr))
}

func (tx Transaction) HasEventLog(eventName string) bool {
	for _, event := range tx.Receipt.EventLogs {
		if event.EventName == eventName {
			return true
		}
	}
	return false
}

func (tx Transaction) GetTransition(transition string) (transitions []Transition) {
	for _, t := range tx.Receipt.Transitions {
		if t.Msg.Tag == transition {
			transitions = append(transitions, t)
		}
	}
	return transitions
}

func (tx Transaction) HasTransition(transition string) bool {
	for _, t := range tx.Receipt.Transitions {
		if t.Msg.Tag == transition {
			return true
		}
	}
	return false
}

func (tx Transaction) IsMint() bool {
	return tx.HasEventLog("MintSuccess") || tx.HasTransition("Mint")
}

func (tx Transaction) IsTransfer() bool {
	return tx.HasTransition("TransferFrom") &&
		tx.GetTransition("TransferFrom")[0].Msg.Params.HasParam("token_id", "Uint256")
}
