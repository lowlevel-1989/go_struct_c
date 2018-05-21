#include <stdio.h>
#include "main.h"

void TestStrucInC(Person *person) {
	printf("Test in C\n");
	printf("Name: %s\n", person->Name);
	printf("Age: %d\n", person->Age);
}

void TestStructReturn() {
	Person person = GetPerson();
	printf("Test in C for Return\n");
	printf("Name: %s\n", person.Name);
	printf("Age: %d\n", person.Age);
}
