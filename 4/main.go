package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.

	fmt.Println("ааа")

	// создание массива
	var numbers [5]int

	// заполнение массива
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}

	// вывод массива
	fmt.Println(numbers)
}
