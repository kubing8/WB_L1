package Solutions

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Вариант с сигналом в выходной канал
func Task6_1() {
	exitCh := make(chan int)
	go func() {
		for {
			select {
			case <-exitCh:
				fmt.Println("Я все..")
				return
			default:
				fmt.Println("Я все еще работаю")
			}
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(7 * time.Second)
	exitCh <- 1
	time.Sleep(1 * time.Second)
}

// Закрытие канала
func Task6_2() {
	exitCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for data := range exitCh {
			fmt.Printf("Получено: %v\n", data)
		}
		wg.Done()
	}()

	exitCh <- 2
	exitCh <- 0

	close(exitCh)
	wg.Wait()
	fmt.Println("Горутина закрыта")
}

// С помощью контекста
func Task6_3() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				fmt.Println("Я работаю")
			}

			time.Sleep(1 * time.Second)
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("Горутина закрыта")
}
