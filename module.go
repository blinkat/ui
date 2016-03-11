package ui

import ()

// var modules = make(map[Handle]Module)

/**
 * module interface
 * all window and module type need fot this
 */
type Module interface {
	Handle() Handle
	GetMessage() bool
	Destory()

	Background() Brush
	SetBackground(b Brush)

	// -- size --
	Size() Size
	SetSize(s Size)

	// pri
	findModule(h Handle) Module

	// ----------- event ----------
	OnClose()
	OnPaint(dc *DeviceContext)
}

// ************ base module **************
type Base struct {
	handle  Handle
	modules map[Handle]Module

	background Brush
}

// get window or module handle
func (b *Base) Handle() Handle {
	return b.handle
}

// destory
func (b *Base) Destory() {
	cdestroyWindow(b.handle)
}

// get window or module message
// if return false this object is quit
func (b *Base) GetMessage() bool {
	return cGetMessage(b.handle)
}

// background
// begin paint will use background brush clear canvas
func (b *Base) Background() Brush {
	return b.background
}

func (b *Base) SetBackground(br Brush) {
	if b.background != nil && b.background != DefaultBackground {
		b.Destory()
	}

	b.background = br
}

/**
 * set a color alpha
 */
func (ba *Base) SetOpacityColor(r, g, b, a uint8) {
	cSetOpacityColor(ba.handle, r, g, b, a)
}

// ---- size ----
func (b *Base) Size() Size {
	return NewSize(cGetSize(b.handle))
}

func (b *Base) SetSize(s Size) {
	cSetSize(b.handle, s.Width, s.Height)
}

// set window opacity
func (b *Base) SetOpacity(a uint8) {
	cSetOpacity(b.handle, a)
}

func (b *Base) Opacity() uint8 {
	return cGetOpacity(b.handle)
}

// ------- events --------
// will close event
func (b *Base) OnClose() {
	b.Destory()
}

func (b *Base) OnPaint(dc *DeviceContext) {
	dc.Clear()
}

// create module base
func NewModuleBase(width, height int, isTool bool, parent Handle) (*Base, error) {
	if !isInit {
		return nil, putError("not init.")
	}

	h, err := ccreateWindow(width, height, isTool, parent)
	if err != nil {
		return nil, err
	}

	r := &Base{
		handle:     h,
		background: DefaultBackground,
		modules:    make(map[Handle]Module),
	}

	return r, nil
}

// private
func (b *Base) findModule(h Handle) Module {
	if v, ok := b.modules[h]; ok {
		return v
	} else {
		for _, v := range b.modules {
			ret := v.findModule(h)
			if ret != nil {
				return ret
			}
		}
	}
	return nil
}
