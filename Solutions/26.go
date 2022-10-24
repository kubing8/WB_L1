package Solutions

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).

Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
func Task26() {
	fmt.Printf("abcd — %v\nabCdefAaf — %v\nAbCda — %v",
		uniqum("abcd"), uniqum("abCdefAaf"), uniqum("AbCda"))
}

// Уникальные ли все символы в слове
func uniqum(str string) bool {
	chars := strings.Split(str, "") // Разделяем строку на символы
	chMap := make(map[string]bool)  // Создаем множество по символам

	// Цикл по всем символам в строке
	for _, char := range chars {
		char = strings.ToLower(char) // Переводим символ в нижний регистр

		// Проверяем есть ли символ в множестве, если есть, возвращаем Fasle
		if _, ok := chMap[char]; ok {
			return false
		}
		chMap[char] = true // Помещаем символ в множество
	}
	return true // Если мы прошли всю строку и не нашли дублирования, возвращаем True
}
