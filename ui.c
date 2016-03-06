#include "inc/ui.h"
#include "stdlib.h"
#include "time.h"

gIcon _DEFAULT_ICON = NULL;
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

// void GenDefaultWinStyle(WinStyle *s)
// {
// 	if (s == NULL) s = malloc(sizeof(WinStyle));
// 	s.x = 0;
// 	s.y = 0;
// 	s.width = 1024;
// 	s.height = 768;
// 	s.icon = _DEFAULT_ICON;
// 	s.title = "unknow window";
// 	s.parent = NULL;
// 	s.style = gWS_DEFAULT;
// }

void gSetDefaultIcon(void* buf)
{
	_DEFAULT_ICON = gLoadIcon(buf, 128, 128);
}