//— Пусть у нас есть интервьюер, который ходит по заданному району и собирает у людей мнения о новом законопроекте. — Интервьюер заранее не знает, сколько людей он повстречает и сколько мнений сможет собрать. Эта случайность иммитируется при помощи рандомно сгенерированного числа, назовём это число N. — Со случайным интервалом (от 300ms до 700ms) интерьвюер будет находить одного из N жителей, получать от него мнение (пусть мнение == случайно сгенерированная строка из 5-ти символов), после чего передавать это мнение в строковый канал. — main горутина должны вычитать все мнения людей, которые собрал интервьюер, после чего завершиться

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func Interview(text chan string) {
	N := rand.Intn(100)
	T := rand.Intn(5)
	T += 3

	for range N + 1 {
		time.Sleep(time.Duration(T*100) * time.Millisecond)
		text <- RandomString(5)

	}

	close(text)
}

func main() {

	text := make(chan string)

	go Interview(text)

	for v := range text {
		fmt.Println(v)
	}
}
