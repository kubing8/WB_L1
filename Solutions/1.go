package Solutions

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	gender string
	age    int
}

func NewHuman(setGender string, setAge int) *Human {
	return &Human{gender: setGender, age: setAge}
}

func (h *Human) ChangeGender(newGender string) {
	h.gender = newGender
}

type Action struct {
	*Human
}

func NewAction(h *Human) Action {
	return Action{h}
}

func (a *Action) GrowUp() {
	a.age++
}

func Task1() {
	// Создаем двух людей me & myFriend:
	me := NewHuman("Male", 21)         // Указатель на структуру Human
	myFriend := NewHuman("Female", 19) // Указатель на структуру Human

	fmt.Printf("Стартовые значения:\n%v\n%v\n", *me, *myFriend)

	// Действия с "me"
	myAction := NewAction(me)       // Назначаем, что действия направлены на "me"
	myAction.ChangeGender("Female") // Меняем гендер на Female у "me"
	myAction.GrowUp()               // У "me" инкрементируем значение возвраста Age
	fmt.Printf("\nИзменения me:\n%v\n", *me)

	friendAction := NewAction(myFriend) // Назначаем, что действия направлены на "myFriend"
	friendAction.ChangeGender("Male")   // Меняем гендер на Male у "myFriend"
	fmt.Printf("\nИзменения myFriend:\n%v\n", *myFriend)
}
