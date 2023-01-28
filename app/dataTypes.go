package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Variables and Functions defined outside any function,
// can be accessed in all other files within the same package
func dataTypes() (string, bool) {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uin8 uint16 uint32 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var age int32
	age = 23
	fmt.Println(age)

	var name string = "Anton"
	fmt.Printf("%T", name) // print type of variable

	const isCool = true // can not reassign

	// shorthand var declaration
	lastName := "unknown"
	lastName = "Kurtin"
	gender, email := "male", "null@null.com"
	fmt.Println(lastName, gender, email)

	// Math
	fmt.Println(math.Floor(2.7)) // 2
	fmt.Println(math.Ceil(2.7))  // 3
	fmt.Println(math.Sqrt(16))   // 4

	// Pointers
	var tickets string
	tickets = "10 tickets"
	fmt.Println(tickets)  // 10 tickets
	fmt.Println(&tickets) // 0xc000108090 (hash which is a memory address for the variable)

	// var prompt string
	// fmt.Scan(&prompt)
	// fmt.Println("You said:", prompt)

	// Arrays
	var names [50]string // should has fixes max size and cant mix types
	names[0] = "Nana"
	names[1] = "Nicole"
	fmt.Println(names[0], names[1])

	// Slice (the same thing but almost always better) (no mixed data types)
	slimes := []string{}
	slimes = append(slimes, "YT", "Gunna", "Lil Durk") // add elements to the end of slice
	fmt.Printf("They are so slimyyy: %v\n", slimes)

	// Loops
	for index, slime := range slimes {
		name := strings.Fields(slime)
		fmt.Println(index, name)
	}

	// Maps (not mixed data types)
	var prettyGril = make(map[string]string)
	prettyGril["firstName"] = "Anna"
	prettyGril["lastName"] = "Guttenberg"
	prettyGril["email"] = "anna-gutten@gmail.com"

	var herAge uint = 24
	prettyGril["age"] = strconv.FormatUint(uint64(herAge), 10)

	// Create slice of maps, slices are dynamic so initial size will be 0
	var girls = make([]map[string]string, 0)
	girls = append(girls, prettyGril)

	// Create new custom type, struct - mixed data type,
	type UserData struct {
		firstName string
		lastName  string
		email     string
		age       uint
	}

	msg := "I am helper function!"
	err := false
	return msg, err
}
