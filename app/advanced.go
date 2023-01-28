package main

import (
	"fmt"
)

// Pointers
//
// - Just like in C, & means "address of" and assigns a pointer which is marked as *T.
// - But * also means "value of", or dereferencing the pointer, and assigns the value.
// - This dual use of * is what gets newbies confused about pointers.

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
	x := 5
	zero(x)             // copy
	fmt.Println("x", x) // x всё еще равен 5

	y := 5
	zeroPointer(&y)     // reference
	fmt.Println("y", y) // y is 0

	// создать экземпляр нового типа
	c := Circle{x: 1, y: 2, z: 3, name: "jd"}
	c.changeName("Depp")

	fmt.Println(c.name, c.x, c.y, c.z)
}
