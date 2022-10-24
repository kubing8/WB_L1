package Solutions

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/
func Task21() {
	me := Person{} // Создание пользователя

	myAuto := Auto{markOfAuto: "Solaris"} // Создание новго авто "Солярис"
	me.toWork(myAuto)                     // Использование нового авто как метода передвижеения

	pubTran := PublicTransport{} // Создание публичного транспорта
	me.toWork(pubTran)           // Использование публичного транспорта как метода передвижеения

	foot := Foot{}            // Созданием нового способа движения: пешком
	adapFoot := Adapter{foot} // Создание адаптера, для связи пользователя с видом передвижения "пешком"
	me.toWork(adapFoot)       // Использование метода "пешком" как способ передвижение
}

// Интерфейс транспорта
type ITransport interface {
	Movement()
}

// Пользователь
type Person struct {
}

// Метод пользователя для передвижения
func (p Person) toWork(transport ITransport) {
	transport.Movement()
}

// Автомобиль
type Auto struct {
	markOfAuto string
}

// Реализация движения для автомобиля
func (a Auto) Movement() {
	fmt.Printf("Я еду на %v\n", a.markOfAuto)
}

// Выбор общественного транспорта
type PublicTransport struct {
}

// Реализация движения для публиного транспорта
func (pt PublicTransport) Movement() {
	fmt.Println("Поеду на метро")
}

// Пойти пешком, но тут нет Movement()
type Foot struct {
}

// Метод пойти пешком
func (f Foot) Move() {
	fmt.Println("Я за экологию - иду пешком")
}

// Адаптер, который позволяет использовать один интерфейс вызова функции как и для транспорта
type Adapter struct {
	Foot
}

// Реализация движения пешком
func (adap Adapter) Movement() {
	adap.Move() // Вызываем метод для "Пеший"
}
