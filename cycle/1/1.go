package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x int
	var score int
	for {
		x = rand.Intn(10)

		if x == 5 {
			fmt.Println("ya tupa loh")
			break
		}

		score += 5
		fmt.Println("ura pobeda, my score:", score)
		time.Sleep(100 * time.Millisecond)

	}

	fmt.Println(score)
}
