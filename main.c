#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "main.h"

void TestStrucInC(Person *person) {
	printf("Test in C\n");
	printf("Name: %s\n", person->Name);
	printf("Age: %d\n", person->Age);

	time_t local_birthday = mktime(gmtime(&person->Birthday));

	printf("[UTC-0] GO Unix Birthday: %d\n", person->Birthday);
	printf("[UTC-LOCAL] GO->C Birthday: %s\n", ctime(&local_birthday));

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

	struct tm tm_birthday = {0, 0, 0, 3, 11-1, 89};

	time_t local_birthday = mktime(&tm_birthday);
	time_t utc_birthday = mktime(gmtime(&local_birthday));
	unsigned int parse_utc0 = local_birthday - utc_birthday;

	time_t birthday = local_birthday + parse_utc0;

	person.Birthday = birthday;

	printf("Test in C for Return\n");
	printf("Name: %s\n", person.Name);
	printf("Age: %d\n", person.Age);
	printf("[UTF-0] C Unix Birthday: %d\n", person.Birthday);
	printf("[UTF-LOCAL] C Birthday: %s\n", ctime(&local_birthday));

	person.Byte = (uint8_t*)(malloc(sizeof(uint8_t)*person.Length));
	memset(person.Byte, 0, person.Length);

	TestPrintByte(person.Byte);
}
