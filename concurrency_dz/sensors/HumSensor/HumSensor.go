package humsensor

import (
	"context"
	"fmt"
	"station/sensors"
	"sync"
)

func PreSensor(ctx context.Context, n int, HuCh chan<- sensors.Measure) {

	fmt.Println("сенсор влажности номер ", n, " начал работу")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("сенсор влажности номер ", n, " завершил работу")
			return
		case HuCh <- sensors.NewMeasure(n, "humidity", sensors.NewData("humidity")):
			fmt.Println("сенсор влажности номер ", n, " передал данные")
			//челик отработал
		}
	}
}

func HuSensorPool(ctx context.Context, HuCh chan sensors.Measure, n int, wg *sync.WaitGroup) {

	for i := 1; i <= n; i++ {
		wg.Go(func() {
			PreSensor(ctx, i, HuCh)
		})
	}
}
