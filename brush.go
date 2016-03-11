package ui

import ()

const (
	BRUSH_SOLID = iota
)

type Brush interface {
	// Handle() BrushHandle
	// Destory()
	Type() int
	genBrush() BrushHandle
	genPen() PenHandle

	Width() int
	SetWidth(i int)
}

type SolidBrush struct {
	r, g, b uint8
	handle  BrushHandle
	width   int // if gen pen.
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

func (s *SolidBrush) Type() int {
	return BRUSH_SOLID
}

func (b *SolidBrush) genBrush() BrushHandle {
	return cCreateSolidBrush(b.r, b.g, b.b)
}

func (s *SolidBrush) genPen() PenHandle {
	return cCreatePen(s.r, s.g, s.b, PEN_SOLID, s.width)
}

func (s *SolidBrush) Width() int {
	return s.width
}

func (s *SolidBrush) SetWidth(i int) {
	s.width = i
}

func NewSolidBrush(r, g, b uint8, w int) Brush {
	return &SolidBrush{
		r:     r,
		g:     g,
		b:     b,
		width: w,
		// handle: cCreateSolidBrush(c.R, c.G, c.B),
	}
}
