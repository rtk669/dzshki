//— Мы с вами пишем универсальное хранилище, основанное на мапе — Ключом в хранилище является строка — Значением в хранилище является тоже строка — Необходимо реализовать это хранилище, допуская вероятность конкурентного использования — Напишите нагрузочные тесты (такие, как мы писали в уроке в теме RWMutex), и сравните время выполнения программы с использованием Mutex и с использованием RWMutex (p.s. не всегда будет прирост в скорости, не нужно этого пугаться)

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	storage := make(map[string]string)

	wg := &sync.WaitGroup{}

	//mtx := sync.Mutex{}

	mtx := sync.RWMutex{}

	x := time.Now()

	for i := range 100 {
		wg.Go(func() {
			for j := range 100 {
				str1 := "aa" + strconv.Itoa(i)
				mtx.Lock()
				storage[str1] = "bb" + strconv.Itoa(j)
				mtx.Unlock()
			}
		})
	}

	for range 100 {
		wg.Go(func() {
			mtx.RLock()
			for k := range storage {
				_ = storage[k]

			}
			mtx.RUnlock()
		})
	}

	wg.Wait()

	fmt.Println("время выполнения ", time.Since(x))
}
