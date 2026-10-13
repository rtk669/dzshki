//— Создать 3 контекста со следующей схемой наследования: 1. parentCtx «— context.Background 2. middleCtx «— parentCtx 3. childCtx «— middleCtx — Создать 3 группы горутин — Завязать каждую группу горутин на определённый контекст — Попробовать разные комбинации порядков отмены контекста, оценить что происходит при этом с каждой отдельной группой горутин

package main

import (
	"context"
	"fmt"
	"time"
)

func Popka(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("popka - ВСЕ :(")
				return
			}
		default:
			fmt.Println("popka дает джазу")

		}
		time.Sleep(300 * time.Millisecond)
	}

}

func Pipka(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("pipka - ВСЕ :(")
				return
			}
		default:
			fmt.Println("pipka дает джазу")

		}
		time.Sleep(600 * time.Millisecond)
	}

}

func Jopka(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("jopka - ВСЕ :(")
				return
			}
		default:
			fmt.Println("jopka дает джазу")

		}
		time.Sleep(900 * time.Millisecond)
	}

}

func main() {
	LContext, LCancel := context.WithCancel(context.Background())
	MContext, MCancel := context.WithCancel(LContext)
	SContext, SCancel := context.WithCancel(MContext)

	go Popka(LContext)

	go Pipka(MContext)

	go Jopka(SContext)

	time.Sleep(2 * time.Second)

	LCancel()

	time.Sleep(2 * time.Second)

	MCancel()

	time.Sleep(2 * time.Second)

	SCancel()

	time.Sleep(7 * time.Second)
}
