package ui

import ()

type DeviceContext struct {
	dc DC
	md Module

	pen         Brush // current brush
	paintBucket Brush
}

func newDeviceContext(dc DC, m Module) *DeviceContext {
	return &DeviceContext{
		dc:          dc,
		md:          m,
		paintBucket: DefaultBackground,
		pen:         DefaultDark,
	}
}

func (d *DeviceContext) SetPaintBucket(b Brush) {
	d.paintBucket = b
}

func (d *DeviceContext) PaintBucket() Brush {
	return d.paintBucket
}

func (d *DeviceContext) SetPen(b Brush) {
	d.pen = b
}

func (d *DeviceContext) Pen() Brush {
	return d.pen
}

func (d *DeviceContext) FillRect(r Rect) {
	br := d.paintBucket.genBrush()
	cFillRect(d.dc, r.Left(), r.Top(), r.Right(), r.Bottom(), br)
	cDestoryBrush(br)
}

func (d *DeviceContext) StrokeRect(r Rect) {
	br := d.pen.genPen()
	cStrokeRect(d.dc, r.Left(), r.Top(), r.Right(), r.Bottom(), br)
	cDestoryPen(br)
}

func (d *DeviceContext) Clear() {
	size := d.md.Size()
	bg := d.md.Background()
	if bg == nil {
		bg = DefaultBackground
	}

	old := d.paintBucket
	d.paintBucket = bg
	d.FillRect(NewRect(0, 0, size.Width, size.Height))
	d.paintBucket = old
}
