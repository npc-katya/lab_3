package mathutils

// функция для вычисления факториала
func Factorial(x int) (y int) {

	y = 1
	for i := 1; i <= x; i++ {
		y = y * i
	}

	return y
}
