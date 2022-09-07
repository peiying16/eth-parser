package storage

import "eth/internal/core"

type IStorage interface {
	Subscribe(string) (bool, error)
	// Unsubscribe(string) (bool, error)
	HasSubscriber(string) (bool, error)

	AddTransation(string, core.Transaction) error
	GetTransactions(string) ([]core.Transaction, error)

	GetBlockNumber() (uint64, error)
	SetBlockNumber(uint64) error
}
