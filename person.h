#include <stdint.h>
#include <time.h>

#define PI 3.14

typedef struct {
	char *Name;
	int Age;
	uint8_t *Byte;
	size_t Length;
	time_t Birthday;
} Person;

extern Person GetPerson();

void TestPrintByte(uint8_t *data);

void TestStrucInC(Person *person);

void TestStructReturn();
