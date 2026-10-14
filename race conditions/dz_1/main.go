package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		// атомики
		var vote atomic.Int64
		wg := &sync.WaitGroup{}

		for range 100 {
			wg.Go(func() {
				for range 1000 {
					vote.Add(1)
				}

			})
		}

		wg.Wait()

		fmt.Println(vote.Load())
	*/

	/*
		// без примитивов

		var vote int

		wg := &sync.WaitGroup{}

		for range 100 {
			wg.Go(func() {
				for range 1000 {
					vote++
				}
			})
		}

		wg.Wait()

		fmt.Println(vote)

	*/

	var vote int

	wg := &sync.WaitGroup{}

	mtx := sync.Mutex{}

	for range 100 {
		wg.Go(func() {
			for range 1000 {
				mtx.Lock()
				vote++
				mtx.Unlock()
			}
		})
	}

	wg.Wait()

	fmt.Println(vote)

}
