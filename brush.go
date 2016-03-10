package ui

import ()

const (
	BRUSH_SOLID = iota
)

type Brush interface {
	Handle() BrushHandle
	Destory()
	Type() int
}

type SolidBrush struct {
	r, g, b uint8
	handle  BrushHandle
}

func (b *SolidBrush) R() uint8 {
	return b.r
}

func (b *SolidBrush) G() uint8 {
	return b.g
}

func (b *SolidBrush) B() uint8 {
	return b.b
}

func (b *SolidBrush) SetR(v uint8) {
	b.r = v
}

func (b *SolidBrush) SetG(v uint8) {
	b.g = v
}

func (b *SolidBrush) SetB(v uint8) {
	b.b = v
}

func (b *SolidBrush) Handle() BrushHandle {
	return b.handle
}

func (b *SolidBrush) Destory() {
	cDestoryBrush(b.handle)
}

func (s *SolidBrush) Type() int {
	return BRUSH_SOLID
}

func CreateSolidBrush(r, g, b uint8) Brush {
	return &SolidBrush{
		r:      r,
		g:      g,
		b:      b,
		handle: cCreateSolidBrush(r, g, b),
	}
	// return solidBrush
}
