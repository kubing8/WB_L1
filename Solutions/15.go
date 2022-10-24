package Solutions

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/

/*
Функция обрезает строку, при условии, что она больше 100 сиволов. Если же строка меньше -> runtime.Error
*/
func Task15() {
	var justString string
	v := createHugeString(1 << 3) // образует: [8 7 6 5 4 3 2 1]
	n := 5                        // сколько первых символов оставит

	someFunc(&justString, v, n)
	fmt.Println(justString)
}

// Возвращает строку обрезанною до n символов
func someFunc(justString, v *string, n int) {
	*justString = string([]rune(*v)[:n])
}

// Создание длинной строки
func createHugeString(num int64) *string {
	var str []string
	for num > 0 {
		sym := string((num % 10) + '0')
		str = append(str, sym)
		num--
	}

	fmt.Println("Nums: ", str)

	ans := strings.Join(str, "")
	return &ans
}
