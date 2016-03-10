package ui

import (
//"fmt"
)

type Window struct {
	*Base
	icon   *Image
	border *Border
}

func CreateDefaultWindow(width, height, px, py int, title string, icon *Image, parent Module) (*Window, error) {
	return createWin(width, height, px, py, title, icon, parent, STYLE_DEFAULT)
}

func CreateCustomWindow(width, height, px, py int, title string, icon *Image, parent Module) (*Window, error) {
	return createWin(width, height, px, py, title, icon, parent, STYLE_CUSTOM)
}

func createWin(width, height, px, py int, title string, icon *Image, parent Module, style int) (*Window, error) {
	b, err := CreateBase(width, height, px, py, title, parent, style)
	if err != nil {
		return nil, err
	}

	w := &Window{
		Base: b,
		icon: icon,
	}
	w.border = NewBorder(w)

	if icon != nil {
		w.SetIcon(w.icon)
	}
	Register(w)

	// // bind event
	w.events.PriEvent.Paint = func(md Module, dc *DeviceContext) bool {
		return false
	}
	w.SetOpacity(0)
	return w, nil
}

// set icon window must be showing
func (w *Window) SetIcon(icon *Image) {
	cSetIcon(w.handle, icon)
}

func (w *Window) Paint(dc *DeviceContext) {
	// // img := NewImageForFile("./test/test.png")
	// dc.Clear(w)
	// w.border.Paint(dc)
}

func (w *Window) Show() {
	w.border.AsynShow()
	w.Base.Show()
}
