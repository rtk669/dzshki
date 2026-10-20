package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(ctx context.Context, transferPoint chan<- int, n int, power int) {
	for {
		// Этот шахтёр завершит своё выполнение СРАЗУ после сигнала о завершении рабочего дня
		fmt.Println("Я шахтёр номер:", n, "Начал добывать уголь!")

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтёр номер:", n, "Добыл уголь:", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтёр номер:", n, "Передал уголь:", power)
		}

	}

}

func MinerPool(ctx context.Context, n int) <-chan int {
	transferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		wg.Go(func() {
			miner(ctx, transferPoint, i, i*10)
		})
	}

	go func() {
		wg.Wait()
		close(transferPoint)
	}()

	return transferPoint
}
