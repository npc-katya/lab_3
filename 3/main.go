package main

import (
	"bufio"
	"fmt"
	"os"
	stringutils "stringutils/package"
)

func main() {

	// Создать пакет stringutils с функцией для переворота строки и использовать его в основной программе.

	fmt.Println("введите текст: ")

	// считывание текста с консоли
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	stringutils.Reverse(text)

	// вывод текста
	fmt.Println(stringutils.Reverse(text))

}
