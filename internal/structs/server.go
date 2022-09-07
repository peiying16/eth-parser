package structs

import "eth/internal/core"

type GetBlockReq struct{}
type GetBlockResp struct {
	Number uint64 `json:"number"`
}

type SubscribeReq struct {
	Address string `json:"address"`
}
type SubscribeResp struct {
	Success bool `json:"success"`
}

type GetTransactionsReq struct {
	Address string `json:"address"`
}

type GetTransactionsResp struct {
	Txs []core.Transaction `json:"txs"`
}
