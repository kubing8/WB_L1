package Solutions

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»)
Символы могут быть unicode.
*/

func Task19() {
	inStr := "日глав本fish語"
	revStr := reversString(&inStr)
	fmt.Printf("%q - %q\n", inStr, revStr)
}

func reversString(str *string) string {
	chars := strings.Split(*str, "")
	var revChars []string
	for i := len(chars) - 1; i >= 0; i-- {
		revChars = append(revChars, chars[i])
	}
	return strings.Join(revChars, "")
}
