package client

import "github.com/neoiss/near-api-go/pkg/types"

type GasPrice struct {
	GasPrice types.Balance `json:"gas_price"`
}
