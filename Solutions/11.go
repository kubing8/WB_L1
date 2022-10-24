package Solutions

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func Task11() {
	firstSet := []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	secondSet := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	resultSet := intersection(&firstSet, &secondSet)
	fmt.Printf("Множество A: %v\nМножество B: %v\nПересечение множест: A и B: %v\n",
		firstSet, secondSet, *resultSet)
}

func intersection(a, b *[]int) *[]int {
	tempMap := make(map[int]bool)

	//Записываем в мапу все первое множество
	for _, val := range *a {
		if _, ok := tempMap[val]; !ok {
			tempMap[val] = true
		}
	}

	res := []int{}

	// В ответ записываем те, которые есть во втором и в первом
	for _, val := range *b {
		if _, ok := tempMap[val]; ok {
			res = append(res, val)
		}
	}

	return &res
}
