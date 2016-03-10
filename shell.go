package ui

/*
// windows option
#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lMsimg32
#include "inc/ui.h"
*/
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"
)

// ========= init ========
func Init() error {
	i := C.gInit()
	if i == 0 {
		return fmt.Errorf("init ui failed")
	}
	return nil
}

// ========= key const ======
const (
	KEY_BACK  = Key(C.gKEY_BACK)
	KEY_TAB   = Key(C.gKEY_TAB)
	KEY_ENTER = Key(C.gKEY_ENTER)
	KEY_SHIFT = Key(C.gKEY_SHIFT)
	KEY_CTRL  = Key(C.gKEY_CTRL)
	KEY_ALT   = Key(C.gKEY_ALT)
	KEY_PAUSE = Key(C.gKEY_PAUSE)
	KEY_CAPS  = Key(C.gKEY_CAPS)
	KEY_ESC   = Key(C.gKEY_ESC)
	KEY_SPACE = Key(C.gKEY_SPACE)

	KEY_PAGEUP   = Key(C.gKEY_PAGEUP)
	KEY_PAGEDOWN = Key(C.gKEY_PAGEDOWN)
	KEY_END      = Key(C.gKEY_END)
	KEY_HOME     = Key(C.gKEY_HOME)

	KEY_LEFT  = Key(C.gKEY_LEFT)
	KEY_UP    = Key(C.gKEY_UP)
	KEY_RIGHT = Key(C.gKEY_RIGHT)
	KEY_DOWN  = Key(C.gKEY_DOWN)

	KEY_PRINT  = Key(C.gKEY_PRINT)
	KEY_INSERT = Key(C.gKEY_INSERT)
	KEY_DELETE = Key(C.gKEY_DELETE)
	KEY_HELP   = Key(C.gKEY_HELP)

	KEY_NUM0 = Key(C.gKEY_NUM0)
	KEY_NUM1 = Key(C.gKEY_NUM1)
	KEY_NUM2 = Key(C.gKEY_NUM2)
	KEY_NUM3 = Key(C.gKEY_NUM3)
	KEY_NUM4 = Key(C.gKEY_NUM4)
	KEY_NUM5 = Key(C.gKEY_NUM5)
	KEY_NUM6 = Key(C.gKEY_NUM6)
	KEY_NUM7 = Key(C.gKEY_NUM7)
	KEY_NUM8 = Key(C.gKEY_NUM8)
	KEY_NUM9 = Key(C.gKEY_NUM9)

	KEY_F1  = Key(C.gKEY_F1)
	KEY_F2  = Key(C.gKEY_F2)
	KEY_F3  = Key(C.gKEY_F3)
	KEY_F4  = Key(C.gKEY_F4)
	KEY_F5  = Key(C.gKEY_F5)
	KEY_F6  = Key(C.gKEY_F6)
	KEY_F7  = Key(C.gKEY_F7)
	KEY_F8  = Key(C.gKEY_F8)
	KEY_F9  = Key(C.gKEY_F9)
	KEY_F10 = Key(C.gKEY_F10)
	KEY_F11 = Key(C.gKEY_F11)
	KEY_F12 = Key(C.gKEY_F12)

	KEY_NUMLOCK = Key(C.gKEY_NUMLOCK)
	KEY_LSHIFT  = Key(C.gKEY_LSHIFT)
	KEY_RSHIFT  = Key(C.gKEY_RSHIFT)
	KEY_LCTRL   = Key(C.gKEY_LCTRL)
	KEY_RCTRL   = Key(C.gKEY_RCTRL)
	KEY_LALT    = Key(C.gKEY_LALT)
	KEY_RALT    = Key(C.gKEY_RALT)

	KEY_A = Key(C.gKEY_A)
	KEY_B = Key(C.gKEY_B)
	KEY_C = Key(C.gKEY_C)
	KEY_D = Key(C.gKEY_D)
	KEY_E = Key(C.gKEY_E)
	KEY_F = Key(C.gKEY_F)
	KEY_G = Key(C.gKEY_G)
	KEY_H = Key(C.gKEY_H)
	KEY_I = Key(C.gKEY_I)
	KEY_J = Key(C.gKEY_J)
	KEY_K = Key(C.gKEY_K)
	KEY_L = Key(C.gKEY_L)
	KEY_M = Key(C.gKEY_M)
	KEY_N = Key(C.gKEY_N)
	KEY_O = Key(C.gKEY_O)
	KEY_P = Key(C.gKEY_P)
	KEY_Q = Key(C.gKEY_Q)
	KEY_R = Key(C.gKEY_R)
	KEY_S = Key(C.gKEY_S)
	KEY_T = Key(C.gKEY_T)
	KEY_U = Key(C.gKEY_U)
	KEY_V = Key(C.gKEY_V)
	KEY_W = Key(C.gKEY_W)
	KEY_X = Key(C.gKEY_X)
	KEY_Y = Key(C.gKEY_Y)
	KEY_Z = Key(C.gKEY_Z)

	KEY_1 = Key(C.gKEY_1)
	KEY_2 = Key(C.gKEY_2)
	KEY_3 = Key(C.gKEY_3)
	KEY_4 = Key(C.gKEY_4)
	KEY_5 = Key(C.gKEY_5)
	KEY_6 = Key(C.gKEY_6)
	KEY_7 = Key(C.gKEY_7)
	KEY_8 = Key(C.gKEY_8)
	KEY_9 = Key(C.gKEY_9)
	KEY_0 = Key(C.gKEY_0)

	KEY_MINUS_UNDERLINE = Key(C.gKEY_MINUS_UNDERLINE) // -_
	KEY_ADDA_ND         = Key(C.gKEY_ADDA_ND)         // +=
	KEY_BACKLASH        = Key(C.gKEY_BACKLASH)        // |\
	KEY_TILDE           = Key(C.gKEY_TILDE)           // `~
	KEY_SEMICONLON      = Key(C.gKEY_SEMICONLON)      // :;
	KEY_QUOTES          = Key(C.gKEY_QUOTES)          // '"
	KEY_COMMA           = Key(C.gKEY_COMMA)           // <,
	KEY_PERIOD          = Key(C.gKEY_PERIOD)          // .>
	KEY_FORWARDSLASH    = Key(C.gKEY_FORWARDSLASH)    // /?
	KEY_BRACKETLEFT     = Key(C.gKEY_BRACKETLEFT)     // [{
	KEY_BRACKETRIGHT    = Key(C.gKEY_BRACKETRIGHT)    // ]}

	KEY_NUM_FORWARDSLASH = Key(C.gKEY_NUM_FORWARDSLASH)
	KEY_NUM_ASTERISK     = Key(C.gKEY_NUM_ASTERISK)
	KEY_NUM_MINUS        = Key(C.gKEY_NUM_MINUS)
	KEY_NUM_ADD          = Key(C.gKEY_NUM_ADD)
	KEY_NUM_DOT          = Key(C.gKEY_NUM_DOT)
)

const (
	STYLE_DEFAULT = C.gWS_DEFAULT
	STYLE_CUSTOM  = C.gWS_CUSTOM
	STYLE_MODULE  = C.gWS_MODULE
)

// ========= type ========
// windows handle
type Handle C.gHANDLE
type DC C.gDC
type Char C.gCHAR
type BrushHandle C.gBrush
type Icon C.gIcon
type PenHandle C.gPen

// ========= function =======
func ccreateWindow(width, height int, title string, px, py, style int, parent Handle) (Handle, error) {
	ct := syscall.StringToUTF16Ptr(title)

	ret := C.gCreateWindow(C.int(width), C.int(height), (*C.wchar_t)(ct), C.int(px), C.int(py), C.int(style), parent)
	if ret == nil {
		return nil, fmt.Errorf("Create window failed.")
	}
	return Handle(ret), nil
}

func cshowWindow(h Handle) {
	C.gShowWindow(h)
}

func cdestroyWindow(h Handle) {
	C.gDestroyWindow(h)
}

func cGetSize(h Handle) (int, int) {
	width := C.int(0)
	height := C.int(0)

	C.gGetSize(h, &width, &height)
	return int(width), int(height)
}

func cSetSize(han Handle, w, h int) {
	C.gSetSize(han, C.int(w), C.int(h))
}

func cSetOpacity(h Handle, a uint8) {
	C.gSetOpacity(h, C.gBYTE(a))
}

func cGetOpacity(h Handle) uint8 {
	return uint8(C.gGetOpacity(h))
}

func cMoveBottom(h Handle) {
	C.gMoveBottom(h)
}

func cMoveTop(h Handle) {
	C.gMoveTop(h)
}

func cGetLocation(h Handle) (int, int) {
	var x, y C.int
	C.gGetLocation(h, &x, &y)
	return int(x), int(y)
}
func cSetLocation(h Handle, x, y int) {
	C.gSetLocation(h, C.int(x), C.int(y))
}

func cCreateSolidBrush(r, g, b uint8) BrushHandle {
	ret := C.gCreateSolidBrush(C.gBYTE(r), C.gBYTE(g), C.gBYTE(b))
	return BrushHandle(ret)
}

func cDestoryBrush(h BrushHandle) {
	C.gDestoryBrush(h)
}

func cSetIcon(h Handle, rgba *Image) {
	C.gSetIcon(h, rgba.Icon)
}

func cLoadIcon(pix []uint8, width, height int) Icon {
	p := unsafe.Pointer(&pix[0])
	return Icon(C.gLoadIcon(p, C.int(width), C.int(height)))
}
func cDestoryIcon(ico Icon) {
	C.gDestoryIcon(ico)
}

func cRePaint(h Handle) {
	C.gRePaint(h)
}

// *************************
// device context
// *************************
func cFillRect(dc DC, left, top, right, bottom int, b BrushHandle) {
	C.gFillRect(dc, C.int(left), C.int(top), C.int(right), C.int(bottom), b)
}

func cClearBackground(hwnd Handle, dc DC) {
	C.gClearBackground(hwnd, dc)
}

func cStrokeRect(dc DC, left, top, right, bottom int, pen PenHandle) {
	C.gStrokeRect(dc, C.int(left), C.int(top), C.int(right), C.int(bottom), pen)
}

// ======= event shell ======
//export gSizeEvent
func gSizeEvent(handle Handle, w, h C.int) {

}

//export gDestoryEvent
func gDestoryEvent(h Handle) {

}

//export gMoveEvent
func gMoveEvent(h Handle, x, y C.int) {
}

//export gFocusEvent
func gFocusEvent(h Handle) {
}

//export gFocusOutEvent
func gFocusOutEvent(h Handle) {
}

//export gPaintEvent
func gPaintEvent(h Handle, dc DC) {
	callModuleEvent(h, _EVENT_PAINT, newDeviceContext(dc))
}

//export gCloseEvent
func gCloseEvent(h Handle) {
	callModuleEvent(h, _EVENT_CLOSE)
}

//export gKeyDownEvent
func gKeyDownEvent(h Handle, k C.int) {
}

//export gKeyUpEvent
func gKeyUpEvent(h Handle, k C.int) {

}

//export gMouseLBDownEvent
func gMouseLBDownEvent(h Handle, x C.int, y C.int) {

}

//export gMouseLBUpEvent
func gMouseLBUpEvent(h Handle, x C.int, y C.int) {

}

//export gMouseLBDoubleEvent
func gMouseLBDoubleEvent(h Handle, x C.int, y C.int) {

}

//export gMouseRBUpEvent
func gMouseRBUpEvent(h Handle, x C.int, y C.int) {

}

//export gMouseRBDownEvent
func gMouseRBDownEvent(h Handle, x C.int, y C.int) {

}

//export gMouseRBDoubleEvent
func gMouseRBDoubleEvent(h Handle, x C.int, y C.int) {

}

//export gMouseMBUpEvent
func gMouseMBUpEvent(h Handle, x C.int, y C.int) {

}

//export gMouseMBDownEvent
func gMouseMBDownEvent(h Handle, x C.int, y C.int) {

}

//export gMouseMBDoubleEvent
func gMouseMBDoubleEvent(h Handle, x C.int, y C.int) {

}

//export gMouseMBWheelEvent
func gMouseMBWheelEvent(h Handle, x C.int, y C.int, wheel C.int) {

}

//export gCreatedEvent
func gCreatedEvent(h Handle) {
	callModuleEvent(h, _EVENT_CREATED)
}

//export gShowEvent
func gShowEvent(h Handle) {
	callModuleEvent(h, _EVENT_SHOW)
}

// ****************** pen ********************************
func cCreatePen(r, g, b uint8, style int, width int) PenHandle {
	return PenHandle(C.gCreatePen(C.int(style), C.gBYTE(r), C.gBYTE(g), C.gBYTE(b), C.int(width)))
}

func cDestoryPen(h PenHandle) {
	C.gDestoryPen(h)
}

const (
	PEN_SOLID = C.gPEN_SOLID
	PEN_DASH  = C.gPEN_DASH
	PEN_DOT   = C.gPEN_DOT
)
