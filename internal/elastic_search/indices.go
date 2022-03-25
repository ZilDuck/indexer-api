package elastic_search

import (
	"fmt"
)

type Indices string

var AuditIndex Indices = "audit"
var NftIndex Indices = "nft"
var NftActionIndex Indices = "nftaction"
var ContractIndex Indices = "contract"
var ContractStateIndex Indices = "contractstate"

// Sets the network and returns the full string
func (i *Indices) Get(network string) string {
	return fmt.Sprintf("zilliqa.%s.%s", network, string(*i))
}

// Sets the network and returns the full string
func (i *Indices) GetByDate(date string) string {
	return fmt.Sprintf("audit-%s", date)
}
