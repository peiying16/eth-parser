package storage

import (
	"eth/internal/core"
	"testing"
)

func TestMemoryStorage_Subscribe(t *testing.T) {
	s := NewMemoryStorage()

	address := "0x161e49d16d5952ddcc38e68b93b02f826138ff4f"
	has, _ := s.HasSubscriber(address)
	if has {
		t.Errorf("HasSubscriber should return false, got %v", has)
	}

	s.Subscribe(address)
	has, _ = s.HasSubscriber(address)
	if !has {
		t.Errorf("HasSubscriber should return true, got %v", has)
	}
}

func TestMemoryStorage_Transation(t *testing.T) {
	s := NewMemoryStorage()

	address := "0x161e49d16d5952ddcc38e68b93b02f826138ff4f"
	s.Subscribe(address)

	s.AddTransation(address, core.Transaction{
		Hash: "0x7a1fda4b79f08a491c82d8ece4f8811a681f40b8ce6fe6d635bbcf20a1b2fd73",
	})
	txs, _ := s.GetTransactions(address)
	if len(txs) != 1 {
		t.Errorf("GetTransactions should return 1 tx, got %v", len(txs))
	}

	s.AddTransation(address, core.Transaction{
		Hash: "0x8a1fda4b79f08a491c82d8ece4f8811a681f40b8ce6fe6d635bbcf20a1b2fd7a",
	})
	txs, _ = s.GetTransactions(address)
	if len(txs) != 2 {
		t.Errorf("GetTransactions should return 2 tx, got %v", len(txs))
	}
}

func TestMemoryStorage_BlockNumber(t *testing.T) {
	s := NewMemoryStorage()

	num := uint64(1234)
	s.SetBlockNumber(num)
	block, _ := s.GetBlockNumber()

	if num != block {
		t.Errorf("GetBlockNumber should return %v tx, got %v", num, block)
	}

}
