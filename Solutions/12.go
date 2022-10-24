package Solutions

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

// Если имеется ввиду "Собственное подмножество", то достаточно вывести одно любое слово
// В данном решение, мы рассматриваем задачу, как получения set этой последовательности

func Task12() {
	words := []string{"cat", "cat", "dog", "cat", "tree", "dog", "-1"}
	set := getSet(words)
	fmt.Printf("Полученные слова: %q\nМножество: %q\n", words, *set)
}

func getSet(words []string) *[]string {
	tempMap := make(map[string]bool)

	// В цикле проверяем, если данной строки еще нет, значит добавляем ее в множество
	for _, word := range words {
		if _, ok := tempMap[word]; !ok {
			tempMap[word] = true
		}
	}

	set := make([]string, len(tempMap))
	ind := 0
	for key, _ := range tempMap {
		set[ind] = key
		ind++
	}
	return &set
}
