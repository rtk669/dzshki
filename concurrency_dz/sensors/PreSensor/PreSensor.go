package presensor

import (
	"context"
	"fmt"
	"station/sensors"
	"sync"
	"time"
)

func PreSensor(ctx context.Context, n int, prCh chan<- sensors.Measure) {

	fmt.Println("сенсор давления номер ", n, " начал работу")
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("сенсор давления номер ", n, " завершил работу")
			return
		case prCh <- sensors.NewMeasure(n, "pressure", sensors.NewData("pressure")):
			fmt.Println("сенсор давления номер ", n, " передал данные")
			//челик отработал
		}
	}
}

func PreSensorPool(ctx context.Context, prCh chan sensors.Measure, n int, wg *sync.WaitGroup) {

	for i := 1; i <= n; i++ {
		wg.Go(func() {
			PreSensor(ctx, i, prCh)
		})
	}
}
