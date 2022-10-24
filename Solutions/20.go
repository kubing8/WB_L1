package Solutions

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func Task20() {
	inStr := "日 глав 本 fish 語"           // Исходная строка
	revStr := reversWordInString(&inStr) // Перевернутые слова в строке
	fmt.Printf("%q - %q\n", inStr, revStr)
}

// Функция переворота слов в строке
func reversWordInString(str *string) string {
	chars := strings.Split(*str, " ") // Делим строку на слова
	var revChars []string             // Объявление результирующего слайса

	// Цикл по всем словам с конца
	for i := len(chars) - 1; i >= 0; i-- {
		revChars = append(revChars, chars[i]) // Добавляем в результирующий слайс слова
	}
	return strings.Join(revChars, " ") // Возвращаем объединение результирующего слайса
}
