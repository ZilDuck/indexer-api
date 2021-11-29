package elastic_cache

import (
	"fmt"
)

type Indices string

var NftIndex Indices = "nft"
var ContractIndex Indices = "contract"

// Sets the network and returns the full string
func (i *Indices) Get(network string) string {
	return fmt.Sprintf("zilliqa.%s.%s", network, string(*i))
}
