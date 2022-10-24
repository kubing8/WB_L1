package Solutions

import (
	"fmt"
	"github.com/xlab/closer"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
func Task4() {
	var workersNum int
	fmt.Println("Введите желаемое кол-во воркеров (>0)")
	fmt.Scan(&workersNum)

	// Создание канала с размером = кол-во воркеров
	dataCh := make(chan interface{}, workersNum)

	exitCh := make(chan int)

	wg := new(sync.WaitGroup) // Создание синхронизирующей группы

	closer.Bind(func() {
		exitCh <- 1
		close(dataCh) // Закрытие канала
		wg.Wait()
		fmt.Println("Канал успешно закрыт")
	})

	// Создаем определенное кол-во горутин (Ворекров)
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done() //

			// можно так же написать через range
			for {
				// Считываем данные из канала
				data, ok := <-dataCh
				// Проверяем, что канал открыт
				if !ok {
					break
				}
				fmt.Printf("Worker %v: %v\n", j, data)
			}
		}(i)
	}

	for i := 0; ; i++ {
		select {
		case <-exitCh:
			closer.Hold()
		default:
			dataCh <- i // Запись в канал
		}
	}

}
