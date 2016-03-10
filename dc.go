package ui

type DeviceContext struct {
	dc DC
}

func newDeviceContext(dc DC) *DeviceContext {
	return &DeviceContext{
		dc: dc,
	}
}

func (d *DeviceContext) FillRect(r Rect, b Brush) {
	cFillRect(d.dc, r.Left(), r.Top(), r.Right(), r.Bottom(), b.Handle())
}

func (d *DeviceContext) StrokeRect(r Rect, p *Pen) {
	cStrokeRect(d.dc, r.Left(), r.Top(), r.Right(), r.Bottom(), p.handle)
}

func (d *DeviceContext) Clear(m Module) {
	cClearBackground(m.Handle(), d.dc)
}
