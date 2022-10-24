package Solutions

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
func Task23() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 8, 9}
	fmt.Printf("До удаления: %v\n", arr)
	delElem(&arr, 5)

	fmt.Printf("После удаления: %v\n", arr)
}

// Какие типы данных мы можем использовать
type TypeNumber interface {
	int | int64 | string | float32
}

// Удаление элемента основанное на перезаписи слайса без одного элемента
func delElem[T TypeNumber](arr *[]T, ind int) {
	*arr = append((*arr)[:ind], (*arr)[ind+1:]...)
}
