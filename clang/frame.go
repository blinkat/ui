package clang

var MainFrame Frame

// all control interface
type Frame interface {
	Handle() Handle
	Destory()
	GetMessage() bool
	Show() // show and updata window

	Size() *Size
	SetSize(s *Size)

	Opacity() uint8
	SetOpacity(a uint8)

	MoveBottom() // move win to layer bottom
	MoveTop()    // move win to layer up

	Location() *Point
	SetLocation(p *Point)
	// modules
	Modules() map[Handle]Frame
	AddModule(m Frame)
	ReomoveModule(m Frame)
	SetModules(ms map[Handle]Frame)

	Rect() *Rectangle
	SetRect(r *Rectangle)

	SetIcon(ico *Image)
	Repaint()

	// ---- event ----
	OnClose() bool     // return true send close to window
	OnPaint(c *Canvas) // begin paint event
	OnSize(s *Size)
	OnDestory()
	OnMove(p *Point)

	OnFocus()
	OnFocusOut()

	OnKeyDown(k Key)
	OnKeyUp(k Key)

	OnMouseDown(p *Point, btn MouseKey)
	OnMouseUp(p *Point, btn MouseKey)
	OnMouseDouble(p *Point, btn MouseKey)
	OnMouseWheel(p *Point, wheel int)

	OnCreated()
	OnShow()

	OnMouseMove(p *Point)
	OnMouseLeave()
	OnMouseHover()

	// pri ------
	mouseFirstEnter() bool
	setMouseFirstEnter(b bool)
}

func findFrameForMain(h Handle) Frame {
	if MainFrame == nil {
		return nil
	}

	if MainFrame.Handle() == h {
		return MainFrame
	} else {
		return findFrame(MainFrame, h)
	}
}

func findFrame(f Frame, h Handle) Frame {
	cls := f.Modules()
	if v, ok := cls[h]; ok {
		return v
	} else if len(cls) > 0 {
		for _, v := range cls {
			r := findFrame(v, h)
			if r != nil {
				return r
			}
		}
	}
	return nil
}
