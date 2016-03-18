package ui

//#include "inc/ui.h"
import "C"

import ()

type RGB struct {
	R, G, B uint8
}

func (r *RGB) genStyle() *cFillStyle {
	col := C.gCreateColor(C.gBYTE(r.R), C.gBYTE(r.G), C.gBYTE(r.B))
	return (*cFillStyle)(C.gCreateSolid(col))
}

func (r *RGB) genColor() *cColor {
	return (*cColor)(C.gCreateColor(C.gBYTE(r.R), C.gBYTE(r.G), C.gBYTE(r.B)))
}

func (r *RGB) Type() FillStyleType {
	return FILL_SOLID
}

func NewRGB(r, g, b uint8) *RGB {
	return &RGB{
		R: r,
		G: g,
		B: b,
	}
}

// c interface
type cColor C.gColor
type cFillStyle C.gFillStyle

type FillStyleType int

const (
	FILL_SOLID  = FillStyleType(0)
	FILL_LINEAR = FillStyleType(1)
)

type FillStyle interface {
	genStyle() *cFillStyle // gen c lang fill style
	genColor() *cColor
	Type() FillStyleType
}
