package core

import (
	"bytes"
	"encoding/json"
	"eth/internal/config"
	utils "eth/internal/utils"
	"io/ioutil"
	"net/http"
)

type Core struct {
	config *config.Config
}

func NewCore(config *config.Config) *Core {
	return &Core{
		config: config,
	}
}

type BlockNumber struct {
	Result string `json:"result"`
}

func (c *Core) GetBlockNumber() (uint64, error) {
	body, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []any{},
		"id":      1,
	})
	resp, err := c.send(body)
	if err != nil {
		return 0, err
	}

	bn := BlockNumber{}
	err = json.Unmarshal(resp, &bn)
	if err != nil {
		return 0, err
	}

	return utils.Hex2int(bn.Result), nil
}

type Block struct {
	Result BlockResult `json:"result"`
}

type BlockResult struct {
	Number       string                   `json:"number"`
	Hash         string                   `json:"hash"`
	LogsBloom    string                   `json:"logsBloom"`
	ExtraData    string                   `json:"extraData"`
	GasLimit     string                   `json:"gasLimit"`
	Transactions []BlockResultTransaction `json:"transactions"`
}

type BlockResultTransaction struct {
	Type             string `json:"type"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
	GasPrice         string `json:"gasPrice"`
	ChainId          string `json:"chainId"`
}

type Transaction struct {
	Type             string `json:"type"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      uint64 `json:"blockNumber"`
	From             string `json:"from"`
	Gas              uint64 `json:"gas"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            uint64 `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex uint64 `json:"transactionIndex"`
	Value            uint64 `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
	GasPrice         uint64 `json:"gasPrice"`
	ChainId          uint64 `json:"chainId"`
}

func (c *Core) GetTransationsByNumber(num uint64) ([]Transaction, error) {
	body, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params": []any{
			utils.Int2hex(num),
			true,
		},
		"id": 83,
	})

	resp, err := c.send(body)
	if err != nil {
		return nil, err
	}

	block := Block{}
	err = json.Unmarshal(resp, &block)
	if err != nil {
		return nil, err
	}

	txs := make([]Transaction, len(block.Result.Transactions))
	for i, tx := range block.Result.Transactions {
		txs[i] = Transaction{
			Type:             tx.Type,
			BlockHash:        tx.BlockHash,
			BlockNumber:      utils.Hex2int(tx.BlockNumber),
			From:             tx.From,
			Gas:              utils.Hex2int(tx.Gas),
			Hash:             tx.Hash,
			Input:            tx.Input,
			Nonce:            utils.Hex2int(tx.Nonce),
			To:               tx.To,
			TransactionIndex: utils.Hex2int(tx.TransactionIndex),
			Value:            utils.Hex2int(tx.Value),
			V:                tx.V,
			R:                tx.R,
			S:                tx.S,
			GasPrice:         utils.Hex2int(tx.GasPrice),
			ChainId:          utils.Hex2int(tx.ChainId),
		}
	}

	return txs, nil
}

func (c *Core) send(body []byte) ([]byte, error) {
	resp, err := http.Post(c.config.Node, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
