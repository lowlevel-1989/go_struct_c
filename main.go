package main

import "fmt"

// #include "main.h"
import "C"

type Person C.Person

//export GetPerson
func GetPerson() C.Person {
	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28

	return ((C.Person)(person))
}

func main() {
	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28

	fmt.Println("Test in Go")
	fmt.Printf("Name: %s\n", C.GoString(person.Name))
	fmt.Printf("Age: %d\n\n", person.Age)

	C.TestStrucInC((*C.Person)(&person))
	fmt.Println()
	C.TestStructReturn()
}
