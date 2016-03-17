package clang

var MainFrame Frame

// all control interface
type Frame interface {
	Handle() Handle
	Destory()
	GetMessage() bool
	Show() // show and updata window

	Size() (int, int)
	SetSize(w, h int)

	Opacity() uint8
	SetOpacity(a uint8)

	MoveBottom() // move win to layer bottom
	MoveTop()    // move win to layer up

	Location() (int, int)
	SetLocation(x, y int)
	// modules
	Modules() map[Handle]Frame
	AddModule(m Frame)
	ReomoveModule(m Frame)
	SetModules(ms map[Handle]Frame)

	PointIn(x, y int) bool // point in window or not

	Repaint()

	// ---- event ----
	OnClose() bool     // return true send close to window
	OnPaint(c *Canvas) // begin paint event
	OnSize(w, h int)
	OnDestory()
	OnMove(x, y int)

	OnFocus()
	OnFocusOut()

	OnKeyDown(k Key)
	OnKeyUp(k Key)

	OnMouseDown(x, y int, btn MouseKey)
	OnMouseUp(x, y int, btn MouseKey)
	OnMouseDouble(x, y int, btn MouseKey)
	OnMouseWheel(x, y, wheel int)

	OnCreated()
	OnShow()

	OnMouseMove(x, y int)
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
