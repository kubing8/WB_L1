package Solutions

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
func Task16() {
	arr := []int{1, 3, 10, 2, 0, 5, 4, 9, 8, 7, 6}
	quick(arr)
	fmt.Printf("New arr: %v\n", arr)
}

// Если длина массива больше 1, то возвращаем сортировку всего массива
func quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return quickAlg(arr, 0, len(arr)-1)
}

func quickAlg(arr []int, beg, end int) []int {
	if beg < end {
		partitionIndex := partition(arr, beg, end) // сортирует элементы: все меньше в лево, больше - в право

		quickAlg(arr, beg, partitionIndex-1) // вызываем рекурсивно функцию с сортировкой левой части
		quickAlg(arr, partitionIndex+1, end) // вызываем рекурсивно функцию с сортировкой правой части
	}

	return arr
}

// Разбивает массив на две части так, чтобы левая часть была меньше конкретного элемента, а другая больше
func partition(arr []int, beg int, end int) int {
	index := beg // начальный индекс ставится начальный

	for i := beg + 1; i <= end; i++ {
		if arr[i] < arr[beg] {
			index++
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
	
	arr[beg], arr[index] = arr[index], arr[beg]
	return index
}
