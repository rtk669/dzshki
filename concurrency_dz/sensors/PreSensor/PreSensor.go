package presensor

import (
	"context"
	"fmt"
	"station/sensors"
	"sync"
)

func PreSensor(ctx context.Context, n int, prCh chan<- sensors.Measure) {

	fmt.Println("сенсор давления номер ", n, " начал работу")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("сенсор давления номер ", n, " завершил работу")
			return
		case prCh <- sensors.NewMeasure(n, "presure", sensors.NewData("presure")):
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
