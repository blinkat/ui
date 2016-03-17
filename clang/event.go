package clang

//#include "inc/ui.h"
import "C"

import (
	"time"
)

type MouseKey int

const (
	MOUSE_LEFT   = MouseKey(0)
	MOUSE_RIGHT  = MouseKey(1)
	MOUSE_MIDDLE = MouseKey(2)
)

//export gSizeEvent
func gSizeEvent(handle Handle, w, h C.int) {
	if r := findFrameForMain(handle); r != nil {
		r.OnSize(int(w), int(h))
	}
}

//export gDestoryEvent
func gDestoryEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		r.OnDestory()
	}
}

//export gMoveEvent
func gMoveEvent(h Handle, x, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMove(int(x), int(y))
	}
}

//export gFocusEvent
func gFocusEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		r.OnFocus()
	}
}

//export gFocusOutEvent
func gFocusOutEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		r.OnFocusOut()
	}
}

//export gPaintEvent
func gPaintEvent(h Handle, dc DC) {
	if r := findFrameForMain(h); r != nil {
		c := NewCanvas(dc, r)
		r.OnPaint(c)
	}
}

//export gCloseEvent
func gCloseEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		if r.OnClose() {
			r.Destory()
		}
	}
}

//export gKeyDownEvent
func gKeyDownEvent(h Handle, k C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnKeyDown(Key(k))
	}
}

//export gKeyUpEvent
func gKeyUpEvent(h Handle, k C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnKeyUp(Key(k))
	}
}

//export gMouseLBDownEvent
func gMouseLBDownEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDown(int(x), int(y), MOUSE_LEFT)
	}
}

//export gMouseLBUpEvent
func gMouseLBUpEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseUp(int(x), int(y), MOUSE_LEFT)
	}
}

//export gMouseLBDoubleEvent
func gMouseLBDoubleEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDouble(int(x), int(y), MOUSE_LEFT)
	}
}

//export gMouseRBUpEvent
func gMouseRBUpEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseUp(int(x), int(y), MOUSE_RIGHT)
	}
}

//export gMouseRBDownEvent
func gMouseRBDownEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDown(int(x), int(y), MOUSE_RIGHT)
	}
}

//export gMouseRBDoubleEvent
func gMouseRBDoubleEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDouble(int(x), int(y), MOUSE_RIGHT)
	}
}

//export gMouseMBUpEvent
func gMouseMBUpEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseUp(int(x), int(y), MOUSE_MIDDLE)
	}
}

//export gMouseMBDownEvent
func gMouseMBDownEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDown(int(x), int(y), MOUSE_MIDDLE)
	}
}

//export gMouseMBDoubleEvent
func gMouseMBDoubleEvent(h Handle, x C.int, y C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseDouble(int(x), int(y), MOUSE_MIDDLE)
	}
}

//export gMouseMBWheelEvent
func gMouseMBWheelEvent(h Handle, x C.int, y C.int, wheel C.int) {
	if r := findFrameForMain(h); r != nil {
		r.OnMouseWheel(int(x), int(y), int(wheel))
	}
}

//export gCreatedEvent
func gCreatedEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		r.OnCreated()
	}
}

//export gShowEvent
func gShowEvent(h Handle) {
	if r := findFrameForMain(h); r != nil {
		r.OnShow()
	}
}

//export gMouseMove
func gMouseMove(h Handle, x, y C.int) {
	if r := findFrameForMain(h); r != nil {
		b := r.mouseFirstEnter()
		if !b {
			r.setMouseFirstEnter(true)
			r.OnMouseHover()
			go hoverTimer(r)
		}

		r.OnMouseMove(int(x), int(y))
	}
}

var hoverPool = make(map[Handle]chan bool)

func hoverTimer(r Frame) {
	sym := make(chan bool, 1)
	hoverPool[r.Handle()] = sym
	for {
		select {
		case <-sym:
			delete(hoverPool, r.Handle())
			r.setMouseFirstEnter(false)
			return
		default:
			px, py := CursorPos()
			if !r.PointIn(px, py) {
				r.OnMouseLeave()
				sym <- false
			} else {
				time.Sleep(10)
			}
		}
	}
}
