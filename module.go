package ui

import "fmt"

type Module interface {
	Handle() Handle
	Show() // show and run loop
	Events() *Caller
	Destory()

	Size() Size
	Opacity() uint8
	SetOpacity(uint8)

	AddModule(Module)
	RemoveModule(Module)
	IndexModule(Module) int

	Paint(*DeviceContext)
	Background() Brush
	SetBackground(br Brush)

	Location() Point
}

// create module
// if combine Base. need call Register function for call event
func CreateBase(width, height, px, py int, title string, parent Module, style int) (*Base, error) {
	m := &Base{}
	m.events = NewCaller(m)

	var ph Handle
	if parent != nil {
		ph = parent.Handle()
	} else {
		ph = nil
	}

	h, err := ccreateWindow(width, height, title, px, py, style, ph)
	if err != nil {
		return nil, err
	}
	m.handle = h
	m.background = DefaultBackground

	m.events.PriEvent.Paint = func(md Module, dc *DeviceContext) bool {
		size := md.Size()
		// pos := md.Location()
		dc.FillRect(NewRect(0, 0, size.Width, size.Height), md.Background())
		md.Paint(dc)

		switch md.(type) {
		case *Window:
			fmt.Println("window")
		case *Border:
			fmt.Println("border")
		}

		return false
	}
	// add def event

	return m, nil
}

// ========== window base ================
// all win module combine this
type Base struct {
	handle     Handle
	modules    []Module
	events     *Caller
	background Brush
}

func (b *Base) Background() Brush {
	return b.background
}

func (b *Base) SetBackground(br Brush) {
	if b.background != nil && b.background != DefaultBackground {
		b.Destory()
	}
	b.background = br
}

func (b *Base) Modules() []Module {
	return b.modules
}

func (b *Base) AddModule(m Module) {
	b.modules = append(b.modules, m)
}

func (b *Base) RemoveModule(m Module) {
	i := b.IndexModule(m)
	if i > 0 {
		b.modules = append(b.modules[:i], b.modules[i+1:]...)
	} else if i == 0 {
		b.modules = b.modules[1:]
	}
}

func (b *Base) IndexModule(m Module) int {
	for k, v := range b.modules {
		if m == v {
			return k
		}
	}
	return -1
}

func (b *Base) Handle() Handle {
	return b.handle
}

func (b *Base) Show() {
	cshowWindow(b.handle)
}

func (b *Base) AsynShow() {
	go b.Show()
}

func (b *Base) AsynShowByChannel(back chan bool) {
	go b.Show()
	back <- true
}

func (b *Base) Events() *Caller {
	return b.events
}

func (b *Base) Destory() {
	cdestroyWindow(b.handle)
}

func (b *Base) Opacity() uint8 {
	return cGetOpacity(b.handle)
}

func (b *Base) SetOpacity(v uint8) {
	cSetOpacity(b.handle, v)
}

func (b *Base) Size() Size {
	width, height := cGetSize(b.handle)
	return NewSize(width, height)
}

func (b *Base) Location() Point {
	x, y := cGetLocation(b.handle)
	return NewPoint(x, y)
}

func (b *Base) RePaint() {
	cRePaint(b.handle)
}

func (b *Base) Paint(dc *DeviceContext) {

}
