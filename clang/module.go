package clang

//#include "inc/ui.h"
import "C"

import (
	"fmt"
)

// ========== Module ==========
type Module struct {
	modules    map[Handle]Frame
	handle     Handle
	mouseFirst bool
}

func (m *Module) Handle() Handle {
	return m.handle
}

func (m *Module) Destory() {
	C.gDestroyWindow(m.handle)
}

func (m *Module) Modules() map[Handle]Frame {
	return m.modules
}

func (m *Module) AddModule(f Frame) {
	m.modules[m.handle] = f
}

func (m *Module) ReomoveModule(f Frame) {
	if _, ok := m.modules[f.Handle()]; ok {
		delete(m.modules, f.Handle())
	}
}

func (m *Module) SetModules(ms map[Handle]Frame) {
	m.modules = ms
}

func (m *Module) GetMessage() bool {
	return int(C.gGetMessage(m.handle)) != 0
}

func (m *Module) Show() {
	C.gShowWindow(m.handle)
}

// ---
func (m *Module) Size() (int, int) {
	width := C.int(0)
	height := C.int(0)

	C.gGetSize(m.handle, &width, &height)
	return int(width), int(height)
}

func (m *Module) SetSize(w, h int) {
	C.gSetSize(m.handle, C.int(w), C.int(h))
}

func (m *Module) SetOpacity(a uint8) {
	C.gSetOpacity(m.handle, C.gBYTE(a))
}

func (m *Module) Opacity() uint8 {
	return uint8(C.gGetOpacity(m.handle))
}

func (m *Module) MoveBottom() {
	C.gMoveBottom(m.handle)
}

func (m *Module) MoveTop() {
	C.gMoveTop(m.handle)
}

func (m *Module) Location() (int, int) {
	var x, y C.int
	C.gGetLocation(m.handle, &x, &y)
	return int(x), int(y)
}

func (m *Module) SetLocation(x, y int) {
	C.gSetLocation(m.handle, C.int(x), C.int(y))
}

func (m *Module) PointIn(x, y int) bool {
	px, py := m.Location()
	w, h := m.Size()

	return (x >= px && y >= py) && (x < px+w && y < py+h)
	// b := (x >= px && y >= py) && (x < px+w && y < py+h)
	// if !b {
	// 	fmt.Println(x, y, px, py, px+w, py+h)
	// }
	// return b
}

func (m *Module) Repaint() {
	C.gRepaint(m.handle)
}

// =============== [ events ] =====================
func (m *Module) OnClose() bool {
	return true
}

func (m *Module) OnPaint(c *Canvas) {
	w, h := m.Size()
	c.FillRect(0, 0, w, h)
}

func (m *Module) OnSize(w, h int) {

}

func (m *Module) OnDestory() {

}

func (m *Module) OnMove(x, y int) {
}

func (m *Module) OnFocus() {
}

func (m *Module) OnFocusOut() {
}

func (m *Module) OnKeyDown(k Key) {
}

func (m *Module) OnKeyUp(k Key) {

}

func (m *Module) OnMouseDown(x, y int, btn MouseKey) {

}

func (m *Module) OnMouseUp(x, y int, btn MouseKey) {

}

func (m *Module) OnMouseDouble(x, y int, btn MouseKey) {

}

func (m *Module) OnMouseWheel(x, y, wheel int) {

}

func (m *Module) OnCreated() {
}

func (m *Module) OnShow() {
}

func (m *Module) OnMouseMove(x, y int) {
}

func (m *Module) OnMouseLeave() {
}

func (m *Module) OnMouseHover() {
}

// ============ private ==========
func (m *Module) mouseFirstEnter() bool {
	return m.mouseFirst
}

func (m *Module) setMouseFirstEnter(b bool) {
	m.mouseFirst = b
}

// ctor
func NewModule(width, height, px, py, style int, parent Frame) (*Module, error) {
	var par Handle
	if parent != nil {
		par = parent.Handle()
	}

	h := C.gCreateWindow(C.int(width), C.int(height), C.int(px), C.int(py), C.int(style), par)
	if h == nil {
		return nil, fmt.Errorf("Create window failed.")
	}

	ret := &Module{
		handle:     Handle(h),
		modules:    make(map[Handle]Frame),
		mouseFirst: false,
	}

	return ret, nil
}
