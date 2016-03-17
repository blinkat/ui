#include "inc/ui.h"
#include <stdlib.h>
#include <time.h>
#include <malloc.h>

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

gColor* gCreateColor(gBYTE r, gBYTE g, gBYTE b)
{
	gColor *ret = (gColor*)malloc(sizeof(gColor));
	ret->R = r;
	ret->G = g;
	ret->B = b;

	return ret;
}

gFillStyle* gCreateLinear(gColor* cols, int* pos, int len)
{
	gFillStyle* ret = (gFillStyle*)malloc(sizeof(gFillStyle));
	ret->Colors = cols;
	ret->Positions = pos;
	ret->Length = len;

	return ret;
}

gFillStyle* gCreateSolid(gColor* c)
{
	gFillStyle* ret = (gFillStyle*)malloc(sizeof(gFillStyle));
	ret->Colors = c;
	ret->Positions = NULL;
	ret->Length = 0;
	return ret;
}

void gDestoryFillStyle(gFillStyle *s)
{
	int i = 0;

	if (s->Length == 0)
	{
		free(s->Colors);
	} 
	else
	{
		for (; i < s->Length; i++)
		{
			free(s->Colors + i);
			if (s->Positions != NULL)
				free(s->Positions + i);
		}
	}


	free(s);
}