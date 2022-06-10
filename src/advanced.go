package main

import (
	"fmt"
)

// Когда мы вызываем функцию с аргументами, аргументы копируются в функцию
func zero(x int) {
	x = 0
}

// Указатели указывают на участок в памяти, где хранится значение
// Используя указатель *int, мы можем изменить значение оригинальной переменной
func zeroPointer(xPtr *int) {
	*xPtr = 0
}

// Структура — это тип, содержащий именованные поля
type Circle struct {
	x, y, z float64
	name    string
}

// Метод — функция особого типа
// Автоматически предоставляет доступ к указателю на Circle для этого метода
func (c *Circle) changeName(name string) {
	c.name = name
}

func advanced() {
	// Pointers
	x := 5
	zero(x)
	fmt.Println("x", x) // x всё еще равен 5

	y := 5
	zeroPointer(&y)
	fmt.Println("y", y) // y is 0

	// создать экземпляр нового типа
	// var c Circle
	// c = Circle{x: 1, y: 2, z: 3, name: "jd"}

	// или
	c := Circle{1, 2, 3, "jd"}
	c.changeName("Depp")

	fmt.Println(c.name, c.x, c.y, c.z)
}
