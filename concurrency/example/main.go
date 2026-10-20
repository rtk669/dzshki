package main

import (
	"context"
	"fmt"
	"practise_example/miner"
	"practise_example/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	PostManContext, PostManCancel := context.WithCancel(context.Background())
	MinerContext, MinerCancel := context.WithCancel(context.Background())

	var coal atomic.Int64
	mtx := sync.Mutex{}
	var mails []string

	minerChan := miner.MinerPool(MinerContext, 3)
	wg := &sync.WaitGroup{}

	wg.Go(func() {
		for v := range minerChan {
			coal.Add(int64(v))
			time.Sleep(1 * time.Second)
		}
	})

	go func() {
		time.Sleep(10 * time.Second)
		MinerCancel()
		PostManCancel()
	}()

	postManChan := postman.PostmanPool(PostManContext, 3)

	wg.Go(func() {
		for v := range postManChan {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()

			time.Sleep(1 * time.Second)
		}
	})

	wg.Wait()

	fmt.Println("добытый уголь: ", coal.Load())
	mtx.Lock()
	fmt.Println("Суммарное количество полученных писем:", len(mails))
	mtx.Unlock()

}
