package Solutions

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/
func Task25_1() {
	fmt.Print("Выведет имя через 3.. 2.. 1.. ")
	Sleep1(3)
	fmt.Println("Ivan")
}

// с помощью time.After || time.Tick
func Sleep1(x int) {
	//<-time.Tick(time.Duration(x) * time.Second)
	<-time.After(time.Second * time.Duration(x)) // Напишет в канал, после прохождения определенного времени
}

func Task25_2() {
	fmt.Print("Выведет имя через 3.. 2.. 1.. ")
	Sleep2(3)
	fmt.Println("Ivan")
}

func Sleep2(x int) {
	startTime := time.Now()

	// Цикл на время
	for time.Since(startTime) < time.Duration(x)*time.Second {
	}
}

