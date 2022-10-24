package Solutions

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
func Task18() {
	var counter Counter    // Создание счетчика
	wg := sync.WaitGroup{} // Создание wait group

	// Цикл на создании 100 горутин
	for i := 0; i < 100; i++ {
		wg.Add(1) // Добавление в группу одного задания
		// Горутина, которая инкрементирует счетчик
		go func() {
			time.Sleep(200 * time.Millisecond)
			counter.Lock()   // блокировка счетчика
			counter.count++  // инкрементирование
			counter.Unlock() // разблокировка счетчика
			wg.Done()        // Выполлнение задания
		}()
	}

	wg.Wait() // Ожидание группы
	fmt.Printf("Счетчик на момент закрытия программы: %v\n", counter.count)
}

// Структура счетчика с возможностью блокировки
type Counter struct {
	sync.Mutex
	count uint
}
