package ui

//#include "inc/ui.h"
import "C"

import (
	// "image"
	// "image/color"
	// "image/draw"
	"math"
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
func (c *Canvas) BeginPath(p *Point) {
	C.gBeginPath(c.handle, C.int(p.X), C.int(p.Y))
}

func (c *Canvas) EndPath() {
	C.gEndPath(c.handle)
}

func (c *Canvas) LineTo(p *Point) {
	C.gLineTo(c.handle, C.int(p.X), C.int(p.Y))
}

func (c *Canvas) LineMoveTo(p *Point) {
	C.gLineMoveTo(c.handle, C.int(p.X), C.int(p.Y))
}

// arc path
// p center of a circle
// r radius
// start angle
// end angle
func (c *Canvas) Arc(p *Point, radius int, start, end float64) {
	rad := 0.0174533
	tsl := start * rad
	tel := end * rad
	r := float64(radius)

	var deg float64
	if r < 5.08 {
		deg = 0.015
	} else if r < 7.62 {
		deg = 0.06
	} else if r < 25.4 {
		deg = 0.075
	} else {
		deg = 0.15
	}

	dte := deg * 25.4 / r

	if tel < tsl {
		tel += 6.28319
	}
	n := int((tel-tsl)/dte + 0.5)
	if n == 0 {
		n = int(6.28319/dte + 0.5)
	}

	ta := tsl
	nx := r * math.Cos(tsl)
	ny := r * math.Sin(tsl)

	c.LineMoveTo(NewPoint(p.X+int(nx), p.Y-int(ny)))
	for i := 1; i < n; i++ {
		ta += dte
		nx = r * math.Cos(ta)
		ny = r * math.Sin(ta)
		c.LineTo(NewPoint(p.X+int(nx), p.Y-int(ny)))
	}
	nx = r * math.Cos(tel)
	ny = r * math.Sin(tel)
	c.LineTo(NewPoint(p.X+int(nx), p.Y-int(ny)))
}

// ********************************
// beizer curse
// ********************************

func quadratic(s, c, p int, t float64) int {
	fs := float64(s)
	fc := float64(c)
	fp := float64(p)
	fr := math.Pow(1.0-t, 2)*fs + (2 * t * (1.0 - t) * fc) + math.Pow(t, 2)*fp
	return int(math.Floor(fr))
}

func cube(s, c1, c2, p int, t float64) int {
	fs := float64(s)
	fc1 := float64(c1)
	fc2 := float64(c2)
	fp := float64(p)
	fr := (fs * math.Pow(1.0-t, 3)) +
		(3.0 * fc1 * t * math.Pow(1.0-t, 2)) +
		(3.0 * fc2 * math.Pow(t, 2) * (1.0 - t)) +
		(fp * math.Pow(t, 3))

	return int(math.Floor(fr))
}

// quadratic beizer
func (c *Canvas) Beizer2(sp, cp, ep *Point) {
	step := 0.001

	for t := 0.0; t <= 1; t += step {
		x := quadratic(sp.X, cp.X, ep.X, t)
		y := quadratic(sp.Y, cp.Y, ep.Y, t)

		c.LineTo(NewPoint(x, y))
	}
}

// cube beizer
func (c *Canvas) Beizer3(sp, cp1, cp2, ep *Point) {
	step := 0.001

	for t := 0.0; t <= 1; t += step {
		x := cube(sp.X, cp1.X, cp2.X, ep.X, t)
		y := cube(sp.Y, cp1.Y, cp2.Y, ep.Y, t)

		c.LineTo(NewPoint(x, y))
	}
}

func (c *Canvas) Fill() {
	s := c.FillStyle.genStyle()
	C.gFill(c.handle, s)
	C.gDestoryFillStyle(s)
}

// ******************************

func (c *Canvas) Stroke() {
	col := c.Line.Color.genColor()
	C.gStroke(c.handle, (*C.struct__gColor)(col), C.int(c.Line.Width))
	C.free(unsafe.Pointer(col))
}

// draw method

func (c *Canvas) Rect(r *Rectangle) {
	c.LineTo(NewPoint(r.X, r.Y))
	c.LineTo(NewPoint(r.Right(), r.Y))
	c.LineTo(NewPoint(r.Right(), r.Bottom()))
	c.LineTo(NewPoint(r.X, r.Bottom()))
	c.LineTo(NewPoint(r.X, r.Y))
}

func (c *Canvas) RoundRect(r *Rectangle, radius int) {
	left, top, right, bottom := r.Left(), r.Top(), r.Right(), r.Bottom()

	c.LineMoveTo(NewPoint(left+radius, r.Y))
	c.LineTo(NewPoint(right-radius+1, r.Y))

	c.LineMoveTo(NewPoint(right, r.Y+radius))
	c.LineTo(NewPoint(right, bottom-radius+1))

	c.LineMoveTo(NewPoint(right-radius, bottom))
	c.LineTo(NewPoint(left+radius-1, bottom))

	c.LineMoveTo(NewPoint(r.X, bottom-radius))
	c.LineTo(NewPoint(r.X, r.Y+radius-1))

	// round
	c.LineMoveTo(NewPoint(r.X+radius, r.Y))
	c.Arc(NewPoint(r.X+radius, r.Y+radius), radius, 90, 180)

	c.LineMoveTo(NewPoint(right, top+radius))
	c.Arc(NewPoint(right-radius, top+radius), radius, 0, 90)

	c.LineMoveTo(NewPoint(right-radius, bottom))
	c.Arc(NewPoint(right-radius, bottom-radius), radius, 270, 360)

	c.LineMoveTo(NewPoint(left, bottom-radius))
	c.Arc(NewPoint(left+radius, bottom-radius), radius, 180, 270)
}

func (c *Canvas) FillRect(r *Rectangle) {
	c.BeginPath(NewPoint(r.X, r.Y))
	c.Rect(r)
	c.EndPath()
	c.Fill()
}

func (c *Canvas) DrawText(str string, pt *Point, size int) {
	p := syscall.StringToUTF16(str)
	f := syscall.StringToUTF16(c.Font.Face)
	col := c.Font.Color.genColor()
	C.gDrawText(c.handle, C.int(pt.X), C.int(pt.Y), C.int(size), (*C.wchar_t)(&p[0]), (*C.struct__gColor)(col), (*C.wchar_t)(&f[0]), C.int(c.Font.Style), C.int(len(p)))
	C.free(unsafe.Pointer(col))
}

// ---------- pix -----------
func (c *Canvas) Pix(p *Point) *RGB {
	var r, g, b C.gBYTE
	C.gGetPix(c.handle, C.int(p.X), C.int(p.Y), &r, &g, &b)

	return NewRGB(uint8(r), uint8(g), uint8(b))
}

func (c *Canvas) SetPix(p *Point, rgb *RGB) {
	C.gSetPix(c.handle, C.int(p.X), C.int(p.Y), C.gBYTE(rgb.R), C.gBYTE(rgb.G), C.gBYTE(rgb.B))
}

func (c *Canvas) DrawImage(img *Image, pt *Point) {
	size := img.Size()

	for x := 0; x < size.Width; x++ {
		for y := 0; y < size.Height; y++ {
			var ret_color RGB
			r, g, b, a := img.RGBAF(x, y)
			if a != 1.0 {
				dc_col := c.Pix(NewPoint(pt.X+x, pt.Y+y))
				dr := float32(dc_col.R) / 255.0
				dg := float32(dc_col.G) / 255.0
				db := float32(dc_col.B) / 255.0

				r = (1-a)*dr + a*r
				g = (1-a)*dg + a*g
				b = (1-a)*db + a*b

				ret_color.R, ret_color.G, ret_color.B, _ = RGBA_F_To_UINT(r, g, b, 0)
			} else {
				ret_color.R, ret_color.G, ret_color.B, _ = img.RGBA(x, y)
			}

			c.SetPix(NewPoint(pt.X+x, pt.Y+y), &ret_color)
		}
	}
}

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
