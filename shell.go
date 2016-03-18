package ui

//#include "inc/ui.h"
import "C"

import (
	"fmt"
	"syscall"
)

var isInit = false

// ========= init ========
func Init() error {
	i := C.gInit()
	if i == 0 {
		return fmt.Errorf("init ui failed")
	}
	isInit = true
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
	STYLE_CHILD   = C.gWS_CHILD
	STYLE_TOOL    = C.gWS_TOOL

	FONT_BOLD      = int(C.gFONT_BOLD)
	FONT_ITALIC    = int(C.gFONT_ITALIC)
	FONT_UNDERLINE = int(C.gFONT_UNDERLINE)
	FONT_STRIKEOUT = int(C.gFONT_STRIKEOUT)
)

// ========= type ========
// windows handle
type Handle C.gHANDLE
type DC C.gDC
type Char C.gCHAR
type cFont C.gFont

func CursorPos() *Point {
	x := C.int(0)
	y := C.int(0)

	C.gGetCursorPos(&x, &y)
	return NewPoint(int(x), int(y))
}

// add font resource
func AddFont(p string) error {
	s := syscall.StringToUTF16(p)
	if C.gAddFontResource((*C.wchar_t)(&s[0])) == 0 {
		return fmt.Errorf("load font faild.", p)
	}
	return nil
}

func RemoveFont(p string) error {
	s := syscall.StringToUTF16(p)
	if C.gRemoveFontResource((*C.wchar_t)(&s[0])) == 0 {
		return fmt.Errorf("remove font faild.", p)
	}
	return nil
}
