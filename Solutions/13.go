package Solutions

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

// Путем использования синтаксиса Go
func Task13_1() {
	var a, b int
	a = 100
	b = -59

	fmt.Printf("Перед заменой:\nA = %v\nB = %v\n", a, b)
	a, b = b, a
	fmt.Printf("\nПосле замены:\nA = %v\nB = %v\n", a, b)
}

// Математически
func Task13_2() {
	var a, b int
	a = 100
	b = -59

	fmt.Printf("Перед заменой:\nA = %v\nB = %v\n", a, b)
	a += b
	b = a - b
	a = a - b
	fmt.Printf("\nПосле замены:\nA = %v\nB = %v\n", a, b)
}
