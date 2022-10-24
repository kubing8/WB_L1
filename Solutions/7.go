package Solutions

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/
func Task7() {
	myMap := make(map[int]int, 1)

	mlock := sync.Mutex{}  // создаем мьютекс
	wg := sync.WaitGroup{} // Создаем wait group

	n := 100
	wg.Add(n) // Добавляем в группу ожиданий n дел
	for i := 0; i < n; i++ {
		// Запускаем горутины, которые инкрементируют одно поле в мапе
		go func(i int) {
			mlock.Lock() // Блокируем
			myMap[0] += 1
			mlock.Unlock() // Разблокировка

			wg.Done() // Выполнение одного задания
		}(i)
	}

	wg.Wait() // Ожидания завершения заданий
	fmt.Printf("Результирующая мапа: %v\n", myMap)
}
