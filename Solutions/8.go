package Solutions

import (
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

/*
	индекс считается с 0 справа

примеры:
01010110 = 86
01010010 = 82

при вводе 86 2 -> получаем 82; и наоборот
*/
func Task8() {
	var num uint64
	var ind int
	fmt.Println("Введте число и индекс для изменения")
	fmt.Scan(&num, &ind)

	var resNum uint64
	mask := 1 << ind
	strBits := strconv.FormatInt(int64(num), 2)

	// Если нужный бит равен 0, то используем ИЛИ, иначе - И
	if strBits[len(strBits)-1-ind] == '0' {
		resNum = num | uint64(mask)
	} else {
		resNum = num & uint64(^mask)
	}

	fmt.Println(resNum)
}
