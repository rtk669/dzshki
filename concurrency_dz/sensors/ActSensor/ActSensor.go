package actsensor

import (
	"context"
	"fmt"
	"station/sensors"
	"sync"
	"time"
)

func ActSensor(ctx context.Context, n int, ActCh chan<- sensors.Measure) {

	fmt.Println("сенсор активности номер ", n, " начал работу")
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("сенсор активности номер ", n, " завершил работу")
			return
		case ActCh <- sensors.NewMeasure(n, "activity", sensors.NewData("activity")):
			fmt.Println("сенсор активности номер ", n, " передал данные")
			//челик отработал
		}
	}
}

func ActSensorPool(ctx context.Context, ActCh chan sensors.Measure, n int, wg *sync.WaitGroup) {

	for i := 1; i <= n; i++ {
		wg.Go(func() {
			ActSensor(ctx, i, ActCh)
		})
	}
}
