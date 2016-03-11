package ui

type WindowShadow struct {
	*Base
	parent Module
	Width  int // border size
}

func (b *WindowShadow) OnPaint(dc *DeviceContext) {
	db := DefaultBorder.(*SolidBrush)
	br := NewSolidBrush(db.r, db.g, db.b, b.Width)
	sz := b.Size()

	dc.SetPen(br)
	dc.StrokeRect(NewRect(b.Width-1, b.Width-1, sz.Width-(b.Width*2), sz.Height-(b.Width*2)))
}

func NewWindowShadow(parent Module, width int) (*WindowShadow, error) {
	sz := parent.Size()
	b, err := NewModuleBase(sz.Width, sz.Height, true, parent.Handle())
	if err != nil {
		return nil, err
	}

	ret := &WindowShadow{
		Base:   b,
		parent: parent,
		Width:  width,
	}
	ret.background = DefaultBorder
	ret.SetOpacity(30)
	ret.SetOpacityColor(0, 0, 0, 0)

	return ret, nil
}
