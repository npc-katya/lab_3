package main

import "fmt"

func main() {

	// Создать срез из массива и выполнить операции добавления и удаления элементов.

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("изначальный массив\n", numbers)

	// добавление элемента
	numbers = append(numbers, 8)
	fmt.Println("после добавления\n", numbers)

	//удаление элемента
	var n = 3
	numbers = append(numbers[:n], numbers[n+1:]...)
	fmt.Println("после удаления\n", numbers)

}
