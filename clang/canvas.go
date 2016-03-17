package clang

//#include "inc/ui.h"
import "C"

import (
	"image"
	"image/color"
	"image/draw"
	"syscall"
	"unsafe"
)

// canvas font

// set canvas font
// new canvas use this font
var DefaultFont = NewFont("宋体", 0, 12)

// ======= canvas =========
type Canvas struct {
	frame  Frame
	handle DC
	Font   *Font

	Line struct {
		Width int
		Color *RGB
	}

	FillStyle FillStyle
}

// begin path
// x, y path start point
func (c *Canvas) BeginPath(x, y int) {
	C.gBeginPath(c.handle, C.int(x), C.int(y))
}

func (c *Canvas) EndPath() {
	C.gEndPath(c.handle)
}

func (c *Canvas) LineTo(x, y int) {
	C.gLineTo(c.handle, C.int(x), C.int(y))
}

func (c *Canvas) LineMoveTo(x, y int) {
	C.gLineMoveTo(c.handle, C.int(x), C.int(y))
}

func (c *Canvas) Stroke() {
	col := c.Line.Color.genColor()
	C.gStroke(c.handle, (*C.struct__gColor)(col), C.int(c.Line.Width))
	C.free(unsafe.Pointer(col))
}

// draw method
func (c *Canvas) FillRect(px, py, width, height int) {
	fs := c.FillStyle.genStyle()
	col := c.Line.Color.genColor()
	C.gFillRect(c.handle, C.int(px), C.int(py), C.int(width), C.int(height), fs, (*C.struct__gColor)(col), C.int(c.Line.Width))
	C.gDestoryFillStyle(fs)
	C.free(unsafe.Pointer(col))
}

func (c *Canvas) FillRoundRect(px, py, width, height, radius int) {
	fs := c.FillStyle.genStyle()
	col := c.Line.Color.genColor()
	C.gFillRoundRect(c.handle, C.int(px), C.int(py), C.int(width), C.int(height), C.int(radius), fs, (*C.struct__gColor)(col), C.int(c.Line.Width))
	C.gDestoryFillStyle(fs)
	C.free(unsafe.Pointer(col))
}

func (c *Canvas) DrawText(str string, px, py, size int) {
	p := syscall.StringToUTF16(str)
	f := syscall.StringToUTF16(c.Font.Face)
	col := c.Font.Color.genColor()
	C.gDrawText(c.handle, C.int(px), C.int(py), C.int(size), (*C.wchar_t)(&p[0]), (*C.struct__gColor)(col), (*C.wchar_t)(&f[0]), C.int(c.Font.Style), C.int(len(p)))
	C.free(unsafe.Pointer(col))
}

// ---------- pix -----------
func (c *Canvas) Pix(x, y int) *RGB {
	var r, g, b C.gBYTE
	C.gGetPix(c.handle, C.int(x), C.int(y), &r, &g, &b)

	return NewRGB(uint8(r), uint8(g), uint8(b))
}

func (c *Canvas) SetPix(x, y int, rgb *RGB) {
	C.gSetPix(c.handle, C.int(x), C.int(y), C.gBYTE(rgb.R), C.gBYTE(rgb.G), C.gBYTE(rgb.B))
}

func (c *Canvas) DrawImage(img image.Image, px, py int) {
	var rgba *image.RGBA
	b := img.Bounds()
	switch img := img.(type) {
	case *image.RGBA:
		rgba = img
	default:
		m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
		rgba = m
	}

	for x := 0; x < b.Dx(); x++ {
		for y := 0; y < b.Dy(); y++ {
			img_col := rgba.At(x, y).(color.RGBA)
			var ret_color RGB
			if img_col.A != 0xff {
				fr := float32(img_col.R) / 255.0
				fg := float32(img_col.G) / 255.0
				fb := float32(img_col.B) / 255.0
				fa := float32(img_col.A) / 255.0

				dc_col := c.Pix(px+x, py+y)
				dr := float32(dc_col.R) / 255.0
				dg := float32(dc_col.G) / 255.0
				db := float32(dc_col.B) / 255.0

				rr := (1-fa)*dr + fa*fr
				rg := (1-fa)*dg + fa*fg
				rb := (1-fa)*db + fa*fb

				ret_color.R = uint8(rr * 0xff)
				ret_color.G = uint8(rg * 0xff)
				ret_color.B = uint8(rb * 0xff)
			} else {
				ret_color.R = img_col.R
				ret_color.G = img_col.G
				ret_color.B = img_col.B
			}

			c.SetPix(px+x, py+y, &ret_color)
		}
	}
}

// func (c *Canvas) DrawImage(img image.Image, px, py int) {
// 	b := img.Bounds()
// 	c.DrawImageSize(img, px, py, b.Dx(), b.Dy())
// }

// canvas ctor
func NewCanvas(h DC, f Frame) *Canvas {
	c := &Canvas{
		frame:  f,
		handle: h,
		Font:   DefaultFont,
	}

	c.Line.Color = DefaultLineColor
	c.Line.Width = 2
	c.FillStyle = DefaultBackground
	return c
}
