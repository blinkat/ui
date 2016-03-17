package clang

//#include "inc/ui.h"
import "C"

import ()

type Font struct {
	Face  string
	Style int
	Color *RGB
	Size  int
}

func NewFont(s string, st, size int) *Font {

	f := &Font{
		Face:  s,
		Style: st,
		Color: DefaultFontColor,
		Size:  size,
	}

	return f
}
