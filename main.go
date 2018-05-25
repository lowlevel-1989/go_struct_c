package main

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <string.h>
// #include "main.h"
import "C"

type Person C.Person

//export GetPerson
func GetPerson() Person {
	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28

	return person
}

func main() {
	b := [4]byte{1, 2, 3, 4}

	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28
	person.Length = (C.size_t)(len(b))
	person.Byte = (*C.uint8_t)(C.malloc( C.sizeof_uint8_t * person.Length ))


	defer C.free(unsafe.Pointer(person.Byte))


	C.memcpy(unsafe.Pointer(person.Byte), unsafe.Pointer(&b), person.Length);

	fmt.Println("Test in Go")
	fmt.Printf("Name: %s\n", C.GoString(person.Name))
	fmt.Printf("Age: %d\n", person.Age)

	C.TestPrintByte( (*C.uint8_t) (&b[0]) )

	fmt.Printf("Bytes: %x\n", b)
	fmt.Printf("DBytes: %x\n", C.GoBytes(unsafe.Pointer(person.Byte), C.int(person.Length)))
	fmt.Printf("LBytes: %d\n\n", person.Length)


	C.TestStrucInC((*C.Person)(&person))
	fmt.Println()

	C.TestStructReturn()
}
