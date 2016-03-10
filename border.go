package ui

import (
	"fmt"
	"image"
	"image/color"
)

const (
	DEF_BORDER_SIZE = 4
	DEF_BORDER_R    = 51
	DEF_BORDER_G    = 122
	DEF_BORDER_B    = 183
)

type Border struct {
	*Base

	// Left        *LeftBorder
	// Right       *RightBorder
	// Top         *TopBorder
	// Bottom      *BottomBorder
	// LeftTop     *LeftTopConverBorder
	// RightTop    *RightTopConverBorder
	// LeftBottom  *LeftBottomConverBorder
	// RightBottom *RightBottomConverBorder
}

func (b *Border) Paint(dc *DeviceContext) {
	fmt.Println("paint")
	pen := NewPen(PEN_SOLID, DEF_BORDER_R, DEF_BORDER_G, DEF_BORDER_B, 1)

	size := b.Size()
	r := NewRect(DEF_BORDER_SIZE-1, DEF_BORDER_SIZE-1, size.Width-(DEF_BORDER_SIZE*2), size.Height-(DEF_BORDER_SIZE*2))
	dc.StrokeRect(r, pen)
}

func NewBorder(m Module) *Border {
	b := &Border{}
	size := m.Size()
	b.Base, _ = CreateBase(size.Width, size.Height, 0, 0, "", m, STYLE_MODULE)
	b.Base.events.module = b
	Register(b)
	return b
}

// ***********************************************************
type LeftBorder struct {
	*Base
	brush Brush
}

type RightBorder struct {
	*Base
}

type TopBorder struct {
	*Base
}

type BottomBorder struct {
	*Base
}

type LeftTopConverBorder struct {
	*Base
}

type RightTopConverBorder struct {
	*Base
}

type LeftBottomConverBorder struct {
	*Base
}

type RightBottomConverBorder struct {
	*Base
}

// **************** gen border image *************
func genTopBorderImage() *image.NRGBA {
	r := image.Rect(0, 0, 1, DEF_BORDER_SIZE)
	ret := image.NewNRGBA(r)

	ret.Set(0, 0, color.NRGBA{DEF_BORDER_R, DEF_BORDER_G, DEF_BORDER_B, 255})
	ret.Set(0, 1, color.NRGBA{DEF_BORDER_R, DEF_BORDER_G, DEF_BORDER_B, 75})
	ret.Set(0, 2, color.NRGBA{DEF_BORDER_R, DEF_BORDER_G, DEF_BORDER_B, 45})
	ret.Set(0, 3, color.NRGBA{DEF_BORDER_R, DEF_BORDER_G, DEF_BORDER_B, 15})

	return ret
}
