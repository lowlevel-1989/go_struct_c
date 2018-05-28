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

// https://golang.org/src/runtime/slice.go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type Person C.Person

//export GetPerson
func GetPerson() Person {
	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28

	return person
}

func main() {
	var b []byte = []byte{1, 2, 3, 4}

	now := time.Now()

	fmt.Println("-- Test date --")
	// utc local
	fmt.Println(now)
	// utc 0
	fmt.Println(now.UTC())
	fmt.Println()
	fmt.Printf("COST C->PI: %v\n\n", C.PI)

	var birthday time.Time = time.Date(
		1989, time.November, 03, 0, 0, 0, 0, time.UTC)

	var person Person
	person.Name = C.CString("Vinicio Valbuena")
	person.Age = 28
	person.Length = (C.size_t)(len(b))
	person.Byte = (*C.uint8_t)(C.malloc( C.sizeof_uint8_t * person.Length ))
	person.Birthday = (C.time_t)(birthday.Unix())

	defer C.free(unsafe.Pointer(person.Byte))

	C.memcpy(unsafe.Pointer(person.Byte), unsafe.Pointer(&b[0]), person.Length);

	// C -> Go slice
	s := slice{
		unsafe.Pointer(person.Byte),
		int(person.Length),
		int(person.Length),
	}

	// ref https://halfrost.com/go_slice/

	data := *(*[]byte)(unsafe.Pointer(&s))

	fmt.Println("C->Go slice")
	fmt.Println("[C] ",person.Byte)
	fmt.Printf("[GO] %p\n", &data[0])
	fmt.Println(data)
	fmt.Println()


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
