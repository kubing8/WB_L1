package Solutions

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/
func Task17() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15}
	index := binSearch(&arr, 13) // Вызов поиска числа в массиве
	fmt.Printf("Индекс числа: %v\n", index)

	fmt.Printf("Индекс числа: %v\n", binSearch(&arr, 12))
	fmt.Printf("Индекс числа: %v\n", binSearch(&arr, 15))
	fmt.Printf("Индекс числа: %v\n", binSearch(&arr, 0))
	fmt.Printf("Индекс числа: %v\n", binSearch(&arr, 6))
	fmt.Printf("Индекс числа: %v\n", binSearch(&arr, -1))
}

// Функция поиска, возвращающая индекс элемента
func binSearch(nums *[]int, num int) int {
	l := 0
	r := len(*nums) - 1

	// Если не в диапозоне -> -1
	if num < (*nums)[l] || num > (*nums)[r] {
		return -1
	}

	// Цикл пока l != r
	for l != r {
		ind := (l + r) / 2 // Устанавливаем индекс на сережину

		// Проверяем, не нашли ли число
		if num == (*nums)[ind] {
			return ind
		}

		// Если число больше середины, рассматриваем правую часть, иначе - левую
		if num > (*nums)[ind] {
			l = ind + 1
		} else {
			r = ind
		}
	}

	// Если индекс не равен числу, значит возвращаем -1
	if (*nums)[r] != num {
		return -1
	}

	// Возвращаем индекс
	return r
}
