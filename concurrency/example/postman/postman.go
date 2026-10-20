package postman

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func postman(ctx context.Context, n int, msgChan chan<- string) {
	for {
		fmt.Println("я почтальон номер ", n, " начал доставку")
		msg := randomize()
		select {
		case <-ctx.Done():
			fmt.Println("я почтальон номер ", n, " завершил работу")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("я почтальон номер ", n, "взял ", msg)
		}

		select {
		case <-ctx.Done():
			fmt.Println("я почтальон номер ", n, " завершил работу")
			return
		case msgChan <- msg:
			fmt.Println("я почтальон номер ", n, "доставил ", msg)
		}

	}
}

func randomize() string {
	randomMsg := rand.Intn(3)

	switch randomMsg {
	case 1:
		{
			return " журнал для огородника"
		}
	case 2:
		{
			return "плакаты с машинками"
		}
	case 0:
		{
			return "повестку"
		}
	default:
		return ""
	}
}

func PostmanPool(ctx context.Context, n int) <-chan string {
	msgChan := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		wg.Go(func() {

			postman(ctx, i, msgChan)
		})
	}

	go func() {
		wg.Wait()
		close(msgChan)
	}()

	return msgChan
}
