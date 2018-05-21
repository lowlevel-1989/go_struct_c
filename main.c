#include "main.h"
#include <stdio.h>

void TestStrucInC(Person *person) {
	printf("Test in C\n");
	printf("Name: %s\n", person->Name);
	printf("Age: %d\n", person->Age);
}
