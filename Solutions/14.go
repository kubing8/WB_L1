package Solutions

import "fmt"

/*
Разработать программу, которая в рантайме способна
определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
func Task14() {
	var X interface{}

	// Можем вывести тип переменной через %T
	X = "it is string"
	fmt.Printf("Тип переменной сейчас: %T\n", X)

	X = 10
	fmt.Printf("Тип переменной сейчас: %T\n", X)

	X = make(chan int)
	fmt.Printf("Тип переменной сейчас: %T\n", X)

	// С помощью fmt.Sprintf может сохранить %T в переменную
	X = true
	str := fmt.Sprintf("%T", X)
	fmt.Printf("Тип переменной сейчас: %q", str)
}
