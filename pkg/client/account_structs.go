package client

import (
	"github.com/neoiss/near-api-go/pkg/types"
	"github.com/neoiss/near-api-go/pkg/types/hash"
)

type AccountView struct {
	Amount        types.Balance      `json:"amount"`
	Locked        types.Balance      `json:"locked"`
	CodeHash      hash.CryptoHash    `json:"code_hash"`
	StorageUsage  types.StorageUsage `json:"storage_usage"`
	StoragePaidAt types.BlockHeight  `json:"storage_paid_at"`

	QueryResponse
}
