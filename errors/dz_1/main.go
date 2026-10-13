package main

import (
	"errors"
	"fmt"
)

func Calculate(x float64, y float64, action string) (float64, error) {
	if (x > 1000 || x < -1000) || (y > 1000 || y < -1000) {
		return 0, errors.New("предельное число")
	}

	switch action {
	case "+":
		{
			z := x + y
			return z, nil
		}
	case "-":
		{
			z := x - y
			return z, nil
		}
	case "/":
		{
			if y == 0 {
				return 0, errors.New("ошибка деления на ноль")
			}
			z := x / y
			return z, nil
		}
	case "*":
		{
			z := x * y
			return z, nil
		}
	}
	return 0, errors.New("не правильный ввод")
}

func main() {

	z, err := Calculate(10, 100, "*")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z)
	}

}
