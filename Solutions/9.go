package Solutions

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
func Task9() {
	n := 10
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	numsCh := make(chan int)

	// Горутина, которая пишет в канал числа
	go func() {
		for _, n := range nums {
			numsCh <- n
		}
		close(numsCh)
	}()

	rea := readCh(numsCh)

	for i := range rea {
		fmt.Print(i, " ")
	}
}

// Возвращает канал, в который записываются числа n*2
func readCh(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}
