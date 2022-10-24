package Solutions

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
func Task24() {
	pointA := newPoint(0, 0) // Создание первой точки
	pointB := newPoint(3, 4) // Создание второй точки

	distance := lenVector(*pointA, *pointB) // Нахождение длины вектора
	fmt.Printf("Между точками: %+v, %+v расстояние = %v", *pointA, *pointB, distance)
}

// Структура точки, содержит поля x, y
type Point struct {
	x, y int
}

// newPoint: конструктор для Point
func newPoint(a, b int) *Point {
	return &Point{a, b} // Возвращает ссылку на новую точку
}

func lenVector(a, b Point) int {
	// Возвращает вычисление длины вектора по формуле из метематики
	return int(math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2)))
}
