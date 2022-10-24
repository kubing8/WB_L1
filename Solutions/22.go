package Solutions

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

// 2^20 = 1 048 576 -- проблемы в +- таких числах не возникнут

func Task22() {
	var a, b big.Int
	a.SetString("102345678901234567890", 10)
	b.SetString("102345678901234567890", 10)

	var mul, div, add, sub big.Int
	mul.Mul(&a, &b)
	div.Div(&a, &b)
	add.Add(&a, &b)
	sub.Sub(&a, &b)

	fmt.Printf("Выполнение операций с числами:\n%v\n%v\n"+
		"Сложение: %v\nВычитание: %v\nУмножением: %v\nДеление: %v\n",
		a.String(), b.String(), add.String(), sub.String(), mul.String(), div.String())

}
