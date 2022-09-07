package storage

import "eth/internal/core"

type IStorage interface {
	// subscriber
	Subscribe(string) (bool, error)
	HasSubscriber(string) (bool, error)

	// transaction
	AddTransation(string, core.Transaction) error
	GetTransactions(string) ([]core.Transaction, error)

	// blocknumber
	GetBlockNumber() (uint64, error)
	SetBlockNumber(uint64) error
}
