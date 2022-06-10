// https://github.com/sefus/the-little-go-book/blob/master/ru/go.md
//
// Go не является объектно-ориентированным языком (ОО-языком).
// В нем нет объектов, нет наследования и многих других понятий, свойственных ОО-языкам,
// полиморфизма или перегрузки.

package main

import (
	"fmt"
)

// Если параметры имеют одинаковый тип, можно использовать короткий синтаксис
func add(a, b int) int {
	return a + b
}

// Структуры
// Точно также, как и с переменными, необъявленные поля содержат нулевые значения.
type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

// Go передает аргументы в функции как копии (by default).
// *X означает указатель на значение типа X
// Очевидно, что копирование указателя будет стоить дешевле с точки зрения ресурсов,
// чем копирование сложной структуры. В 64 битной системе указатель занимает 64 бита памяти.
// Если мы имеем структуру со множеством полей, то создание копии будет дорогой операцией.
// Смысл указателей в том, что они дают общий доступ к значениям.
// Все это не означает, что вам всегда нужно использовать указатели.
func Super(s *Saiyan) {
	s.Power += 10000
}

// Мы можем ассоциировать метод со структурой
func (s *Saiyan) ChangeName(newName string) {
	s.Person.Name = newName
}

func theLittleGoBook() (bool, string) {
	// По умолчанию, Go присваивает переменным нулевые значения.
	// Для целых чисел это 0, для булевых false, для строк "" и так далее
	var power int
	power = 9000
	// var power int = 9000
	// power := 9000
	fmt.Printf("It's over %d\n", power)

	goku := Saiyan{
		Person: &Person{},
		Power:  9001,
	}

	// goku := Saiyan{}
	// goku := Saiyan{Name: "Goku"}
	// goku.Power = 9000
	// goku := Saiyan{"Goku", 9000}

	Super(&goku) // мы использовали оператор & для получения адреса
	goku.ChangeName("Goku")
	fmt.Println(goku)
	goku.Introduce()

	// fmt.Println(goku.Name)
	// fmt.Println(goku.Person.Name)

	// в Go есть встроенная функция new, которая используется для выделения памяти
	// "goku := new(Saiyan)" the same as "goku := &Saiyan{}"

	// Поля могут быть любого типа, включая другие структуры и типы
	// gohan := &Saiyan{
	//     Name: "Gohan",
	//     Power: 1000,
	//     Father: &Saiyan {
	//         Name: "Goku",
	//         Power: 9001,
	//         Father: nil,
	//     },
	// }

	// Предпочтительным способом обработки ошибок в Go является возвращение значений
	// вместо исключений.

	// Инициализация в условии
	// if err := process(); err != nil {
	//     return err
	// }

	// пустой интерфейс без методов: interface{}.
	// Так как каждый тип реализует все 0 методов этого интерфейса, то можно сказать,
	// что в неявном виде каждый тип реализует пустой интерфейс

	// Для преобразования переменной в определенный тип, используйте .(ТИП)
	// a.(int) + b.(int)

	// Строки и массивы байтов тесто связаны. Мы можем легко конвертировать одно в другое:
	// Заметьте, что когда вы используете []byte(X) или string(X), вы создаёте копию данных.
	// Это необходимо, так как строки в Go неизменяемы.
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)
	fmt.Println(strb, byts)

	// Тип функция
	type Add func(a int, b int) int

	return false, "\nhello from the little go book ;)"
}
