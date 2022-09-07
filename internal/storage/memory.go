package storage

import (
	"errors"
	"eth/internal/core"
	"sync"
)

type MemoryStorage struct {
	transations map[string]map[string]core.Transaction
	lock        sync.RWMutex

	block     uint64
	blockLock sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		transations: make(map[string]map[string]core.Transaction),
	}
}

func (m *MemoryStorage) Subscribe(address string) (bool, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.transations[address]; ok {
		return false, nil
	}
	m.transations[address] = make(map[string]core.Transaction)
	return true, nil
}

// func (m *MemoryStorage) Unsubscribe(address string) (bool, error) {
// 	m.lock.Lock()
// 	defer m.lock.Unlock()
// 	if _, ok := m.transations[address]; ok {
// 		return false, nil
// 	}
// 	delete(m.transations, address)
// 	return true, nil
// }

func (m *MemoryStorage) HasSubscriber(address string) (bool, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.transations[address]
	return ok, nil
}

func (m *MemoryStorage) AddTransation(subscriber string, t core.Transaction) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.transations[subscriber]; ok {
		m.transations[subscriber][t.Hash] = t
	}
	return nil
}

func (m *MemoryStorage) GetTransactions(address string) ([]core.Transaction, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.transations[address]; ok {
		txs := make([]core.Transaction, len(m.transations[address]))
		i := 0
		for _, v := range m.transations[address] {
			txs[i] = v
			i++
		}
		return txs, nil
	}
	return make([]core.Transaction, 0), nil
}

func (m *MemoryStorage) SetBlockNumber(num uint64) error {
	m.blockLock.Lock()
	defer m.blockLock.Unlock()
	if m.block != 0 && m.block+1 != num {
		return errors.New("invalid block number")
	}
	m.block = num
	return nil
}

func (m *MemoryStorage) GetBlockNumber() (uint64, error) {
	m.blockLock.RLock()
	defer m.blockLock.RUnlock()
	return m.block, nil
}
