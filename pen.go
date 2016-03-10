package ui

type Pen struct {
	handle PenHandle
	width  int

	r, g, b uint8
}

func (p *Pen) Destory() {
	cDestoryPen(p.handle)
}

func NewPen(style int, r, g, b uint8, width int) *Pen {
	p := &Pen{
		r:      r,
		g:      g,
		b:      b,
		width:  width,
		handle: cCreatePen(r, g, b, style, width),
	}
	return p
}
