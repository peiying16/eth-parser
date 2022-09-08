package main

import (
	"eth/cmd/routers"
	"eth/internal/config"
	"eth/internal/core"
	"eth/internal/services"
	"eth/internal/storage"
	"fmt"
	"time"
)

func main() {
	cfg := config.NewConfig()
	c := core.NewCore(cfg)
	s := storage.NewMemoryStorage()
	// scan(c, s)

	ch := make(chan uint64) // simulate message queue
	go sendBlockToMQ(ch, c, s)
	go subscribeToMQ(ch, c, s)

	svc := services.NewTransationSvc(s)

	_, err := routers.NewServer(":8080", svc)
	if err != nil {
		panic(err)
	}
}

// func scan(c *core.Core, s storage.IStorage) {
// 	ticker := time.NewTicker(time.Second * 5)

// 	for {
// 		<-ticker.C
// 		bs, err := s.GetBlockNumber()
// 		if err != nil {
// 			fmt.Println("storage GetBlockNumber error", err)
// 			continue
// 		}
// 		cs, err := c.GetBlockNumber()
// 		if err != nil {
// 			fmt.Println("core GetBlockNumber error", err)
// 			continue
// 		}
// 		if cs > bs { // compare both bs and cs
// 			num := bs + 1
// 			err := processBlock(c, s, num)
// 			if err != nil {
// 				s.SetBlockNumber(num)
// 			}
// 		}
// 	}
// }

// func processBlock(c *core.Core, s storage.IStorage, blockNumber uint64) error {
// 	txs, err := c.GetTransationsByNumber(blockNumber)
// 	if err != nil {
// 		fmt.Println("GetTransationsByNumber error", err)
// 		return err
// 	}
// 	for _, tx := range txs {
// 		// eth
// 		from := tx.From
// 		hasSub, err := s.HasSubscriber(from)
// 		if err != nil {
// 			return err
// 		}
// 		if hasSub {
// 			s.AddTransation(from, tx)
// 		}

// 		to := tx.To
// 		hasSub, err = s.HasSubscriber(to)
// 		if err != nil {
// 			return err
// 		}
// 		if hasSub {
// 			s.AddTransation(to, tx)
// 		}

// 		// erc-20 TODO
// 	}
// 	return nil
// }

// // simulate
// // every 5 seconds, check the current block number from chain
// // compare current block number from storage
// // if current block number is less than the one on chain
// // push to message queue
func sendBlockToMQ(ch chan uint64, c *core.Core, s storage.IStorage) {
	ticker := time.NewTicker(time.Second * 5)
	start := uint64(15486257)

	for {
		<-ticker.C
		cs, err := c.GetBlockNumber()
		if err != nil {
			fmt.Println("core GetBlockNumber error", err)
			continue
		}
		if cs > start { // compare both start and cs
			start++
			ch <- start // push to message queue
		}
	}
}

// // subscribe to mq
// // if receive a block number
// // get a block and loop for all transations
// // if the to or from address matches the subscriber
// // add to storage
func subscribeToMQ(ch chan uint64, c *core.Core, s storage.IStorage) {
	for {
		num := <-ch

		fmt.Printf("received block number %d\n", num)

		txs, err := c.GetTransationsByNumber(num)
		if err != nil {
			fmt.Println("GetTransationsByNumber error", err)
			continue
		}
		for _, tx := range txs {
			from := tx.From
			hasSub, err := s.HasSubscriber(from)
			if err != nil {
				fmt.Println("HasSubscriber error", err)
			}
			if hasSub {
				err = s.AddTransation(from, tx)
				if err != nil {
					fmt.Println("AddTransation error", err)
				}
			}

			to := tx.To
			hasSub, err = s.HasSubscriber(to)
			if err != nil {
				fmt.Println("HasSubscriber error", err)
			}
			if hasSub {
				err = s.AddTransation(to, tx)
				if err != nil {
					fmt.Println("AddTransation error", err)
				}
			}
		}

		err = s.SetBlockNumber(num)
		if err != nil {
			fmt.Println("SetBlockNumber error", err)
		}
		fmt.Printf("finished block number %d\n", num)
	}
}
