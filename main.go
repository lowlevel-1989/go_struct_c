package main

import (
	"fmt"
	"unsafe"
	"time"
)

// #include <stdlib.h>
// #include <time.h>
// #include <string.h>
// #include "person.h"
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
	var b [4]byte = [4]byte{1, 2, 3, 4}

	now := time.Now()

	fmt.Println("-- Test date --")
	// utc local
	fmt.Println(now)
	// utc 0
	fmt.Println(now.UTC())
	fmt.Println()

	var birthday time.Time = time.Date(
		1989, time.November, 03, 0, 0, 0, 0, time.UTC)

	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28
	person.Length = (C.size_t)(len(b))
	person.Byte = (*C.uint8_t)(C.malloc( C.sizeof_uint8_t * person.Length ))
	person.Birthday = (C.time_t)(birthday.Unix())

	defer C.free(unsafe.Pointer(person.Byte))

	C.memcpy(unsafe.Pointer(person.Byte), unsafe.Pointer(&b), person.Length);

	fmt.Println("Test in Go")
	fmt.Printf("Name: %s\n", C.GoString(person.Name))
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("[UTC-0] GO Unix Birthday: %d\n", time.Unix(int64(person.Birthday), 0).UTC().Unix())
	fmt.Printf("[UTC-0] GO Birthday: %s\n", time.Unix(int64(person.Birthday), 0).UTC())

	C.TestPrintByte( (*C.uint8_t) (&b[0]) )

	fmt.Printf("Bytes: %x\n", b)
	fmt.Printf("DBytes: %x\n", C.GoBytes(unsafe.Pointer(person.Byte), C.int(person.Length)))
	fmt.Printf("LBytes: %d\n\n", person.Length)


	C.TestStrucInC((*C.Person)(&person))
	fmt.Println()

	C.TestStructReturn()
}
