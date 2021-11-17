package elastic_cache

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/config"
)

type Indices string

var NftIndex Indices = "nft"
var TransactionIndex Indices = "transaction"

// Sets the network and returns the full string
func (i *Indices) Get() string {
	return fmt.Sprintf("%s.%s.%s", config.Get().Network, config.Get().Index, string(*i))
}
