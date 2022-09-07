package services

import (
	"eth/internal/core"
	"eth/internal/storage"
)

type TransationSvc struct {
	s storage.IStorage
}

func NewTransationSvc(s storage.IStorage) *TransationSvc {
	return &TransationSvc{
		s: s,
	}
}

func (t *TransationSvc) GetCurrentBlock() (uint64, error) {
	return t.s.GetBlockNumber()
}

func (t *TransationSvc) Subscribe(address string) (bool, error) {
	return t.s.Subscribe(address)
}

func (t *TransationSvc) GetTransactions(address string) ([]core.Transaction, error) {
	return t.s.GetTransactions(address)
}
