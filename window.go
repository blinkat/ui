package ui

import (
//"fmt"
)

var mainModule Module

func findModule(h Handle) Module {
	if mainModule == nil {
		return nil
	}

	if mainModule.Handle() == h {
		return mainModule
	}

	r := mainModule.findModule(h)
	if r == nil {
		log("cannot found module", _LOG_ERR)
	}
	return r
}

// run main loop
func RunMainLoop(w Module) {
	mainModule = w
	for w.GetMessage() {
	}
}

// window type
type Window struct {
	*Base

	Shadow     *WindowShadow
	borderSize int
}

func (w *Window) OnPaint(dc *DeviceContext) {
	// dc.Clear()
	size := w.Size()

	dark := DefaultDark.(*SolidBrush)
	dc.SetPaintBucket(dark)
	dc.FillRect(NewRect(0, 0, size.Width, size.Height))

	r := NewRect(w.borderSize-1, w.borderSize-1, size.Width-(w.borderSize*2), size.Height-(w.borderSize*2))
	dc.SetPaintBucket(w.background)
	dc.FillRect(NewRect(r.X+1, r.Y+1, r.Width-1, r.Height-1))

	dc.SetPen(DefaultBorder)
	dc.StrokeRect(r)

	w.SetOpacityColor(dark.r, dark.g, dark.b, 0)
}

func (w *Window) findModule(h Handle) Module {
	if w.Shadow.handle == h {
		return w.Shadow
	}
	return w.Base.findModule(h)
}

func NewWindow(w, h int, parent Handle) (*Window, error) {
	b, err := NewModuleBase(w, h, false, parent)
	if err != nil {
		return nil, err
	}

	win := &Window{
		Base:       b,
		borderSize: 4,
	}

	win.Shadow, err = NewWindowShadow(win, win.borderSize)
	if err != nil {
		win.Destory()
		return nil, err
	}

	// win.SetOpacityColor(win.background., g, b, a)
	return win, nil
}
