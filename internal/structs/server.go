package structs

import "eth/internal/core"

// block
type GetBlockReq struct{}
type GetBlockResp struct {
	Number uint64 `json:"number"`
}

// subscribe
type SubscribeReq struct {
	Address string `json:"address"`
}
type SubscribeResp struct {
	Success bool `json:"success"`
}

// transactions
type GetTransactionsReq struct {
	Address string `json:"address"`
}
type GetTransactionsResp struct {
	Txs []core.Transaction `json:"txs"`
}
