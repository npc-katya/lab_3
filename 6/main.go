package main

import "fmt"

func main() {

	// Написать программу, которая создает срез из строк и находит самую длинную строку.

	// создание и заполнение среза строк
	str := []string{"aa", "aaa", "a"}

	k := 0
	m := 0
	j := 0

	// цикл, который находит самую длинную строку
	for i := 0; i < len(str); i++ {
		m = len(str[i])
		if k < m {
			k = m
			j = i
		}
	}

	// вывод
	fmt.Println("все строки:\n", str)
	fmt.Println("самая длинная строка:\n", str[j])

}
