package Solutions

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func Task5() {
	var workTime int64 // Время работы канала
	workTime = 2
	wg := sync.WaitGroup{}
	dataCh := make(chan interface{})

	wg.Add(1)
	go listenChan(&wg, dataCh)

	startTime := time.Now()
	i := 0

	// Цикл, который работает определенное время
	for time.Since(startTime) < time.Duration(workTime)*time.Second {
		dataCh <- i
		i++
	}
	// Если надо удостовериться в работоспособности:
	//fmt.Println("Time work: ", time.Now().Sub(startTime))
	close(dataCh)

	wg.Wait()
	fmt.Println("Завершено")

}

func listenChan(wg *sync.WaitGroup, c chan interface{}) {
	for data := range c {
		fmt.Printf("Получено: %v\n", data)
	}
	wg.Done()
}
