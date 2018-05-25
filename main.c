#include <stdio.h>
#include <stdlib.h>
#include "main.h"

void TestStrucInC(Person *person) {
	printf("Test in C\n");
	printf("Name: %s\n", person->Name);
	printf("Age: %d\n", person->Age);

	TestPrintByte(person->Byte);
}

void TestPrintByte(uint8_t *data){
	printf("[ ");
	for ( int i=0; i<4; i++ )
		printf("%x ", data[i]);
	printf("]\n\n");
}

void TestStructReturn() {
	Person person = GetPerson();
	person.Length = 4;

	printf("Test in C for Return\n");
	printf("Name: %s\n", person.Name);
	printf("Age: %d\n", person.Age);

	person.Byte = (uint8_t*)(malloc(sizeof(uint8_t)*person.Length));

	TestPrintByte(person.Byte);
}
