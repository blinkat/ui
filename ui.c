#include "inc/ui.h"
#include "stdlib.h"
#include "time.h"

void newGUID(gCHAR str)
{
	srand(time(NULL));
	swprintf(str, GUID_LENGTH, L"{%08X-%04X-%04X-%04X-%04X%04X%04X}",
	         rand() & 0xffffffff,
	         rand() & 0xffff,
	         rand() & 0xffff,
	         rand() & 0xffff,
	         rand() & 0xffff, rand() & 0xffff, rand() & 0xffff
	        );
	str[GUID_LENGTH] = '\0';
}

int gInit()
{
	return 1;
}