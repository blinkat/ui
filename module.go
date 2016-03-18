package ui

//#include "inc/ui.h"
import "C"

import (
	"fmt"
	"unsafe"
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
func (m *Module) Size() *Size {
	width := C.int(0)
	height := C.int(0)

	C.gGetSize(m.handle, &width, &height)
	return NewSize(int(width), int(height))
}

func (m *Module) SetSize(p *Size) {
	C.gSetSize(m.handle, C.int(p.Width), C.int(p.Height))
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

func (m *Module) Location() *Point {
	var x, y C.int
	C.gGetLocation(m.handle, &x, &y)
	return NewPoint(int(x), int(y))
}

func (m *Module) SetLocation(p *Point) {
	C.gSetLocation(m.handle, C.int(p.X), C.int(p.Y))
}

func (m *Module) Rect() *Rectangle {
	var l, t, r, b C.int
	C.gGetRect(m.handle, &l, &t, &r, &b)
	return Rect(int(l), int(t), int(r), int(b))
}

func (m *Module) SetRect(r *Rectangle) {
	C.gSetRect(m.handle, C.int(r.X), C.int(r.Y), C.int(r.Width), C.int(r.Height))
}

func (m *Module) Repaint() {
	C.gRepaint(m.handle)
}

func (m *Module) SetIcon(img *Image) {
	size := img.Size()
	pixs := img.src.Pix
	C.gSetIcon(m.handle, unsafe.Pointer(&pixs[0]), C.int(size.Width), C.int(size.Height))
}

// =============== [ events ] =====================
func (m *Module) OnClose() bool {
	return true
}

func (m *Module) OnPaint(c *Canvas) {
	s := m.Size()
	c.FillRect(NewRect(0, 0, s.Width, s.Height))
}

func (m *Module) OnSize(size *Size) {

}

func (m *Module) OnDestory() {

}

func (m *Module) OnMove(p *Point) {
}

func (m *Module) OnFocus() {
}

func (m *Module) OnFocusOut() {
}

func (m *Module) OnKeyDown(k Key) {
}

func (m *Module) OnKeyUp(k Key) {

}

func (m *Module) OnMouseDown(p *Point, btn MouseKey) {

}

func (m *Module) OnMouseUp(p *Point, btn MouseKey) {

}

func (m *Module) OnMouseDouble(p *Point, btn MouseKey) {

}

func (m *Module) OnMouseWheel(p *Point, wheel int) {

}

func (m *Module) OnCreated() {
}

func (m *Module) OnShow() {
}

func (m *Module) OnMouseMove(p *Point) {
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
func NewModule(size *Size, p *Point, style int, parent Frame) (*Module, error) {
	var par Handle
	if parent != nil {
		par = parent.Handle()
	}

	h := C.gCreateWindow(C.int(size.Width), C.int(size.Height), C.int(p.X), C.int(p.Y), C.int(style), par)
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
