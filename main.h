#include <stdint.h>

typedef struct {
	char *Name;
	int Age;
	uint8_t *Byte;
	uint32_t Length;
} Person;

extern Person GetPerson();

void TestPrintByte(uint8_t *data);

void TestStrucInC(Person *person);

void TestStructReturn();
