package main

import (
	"fmt"
	mathutils "mathutils/package"
	"os"
)

func main() {
	// Создать пакет mathutils с функцией для вычисления факториала числа.

	// создание переменной
	var x, y int

	// получение значения переменной
	fmt.Print("введите число: ")
	fmt.Fscan(os.Stdin, &x)

	// вызов пакета mathutils с функцией для вычисления факториала числа.
	y = mathutils.Factorial(x)
	fmt.Println(y)
}
