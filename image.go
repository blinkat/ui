package ui

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Image struct {
	src *image.RGBA
}

func (i *Image) At(x, y int) color.RGBA {
	return i.src.At(x, y).(color.RGBA)
}

func (i *Image) RGBA(x, y int) (r uint8, g uint8, b uint8, a uint8) {
	c := i.At(x, y)
	r = c.R
	g = c.G
	b = c.B
	a = c.A
	return
}

func (i *Image) RGBAF(x, y int) (r float32, g float32, b float32, a float32) {
	c := i.At(x, y)
	r = float32(c.R) / 255.0
	g = float32(c.G) / 255.0
	b = float32(c.B) / 255.0
	a = float32(c.A) / 255.0
	return
}

func (i *Image) Size() *Size {
	b := i.src.Bounds()
	return NewSize(b.Dx(), b.Dy())
}

func RGBA_F_To_UINT(r, g, b, a float32) (uint8, uint8, uint8, uint8) {
	return uint8(r * 0xff),
		uint8(g * 0xff),
		uint8(b * 0xff),
		uint8(a * 0xff)
}

func LoadImage(path string) (*Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return NewImage(img), nil
}

func NewImage(img image.Image) *Image {
	b := img.Bounds()

	var rgb *image.RGBA
	switch img := img.(type) {
	case *image.RGBA:
		rgb = img
	default:
		m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
		rgb = m
	}

	return &Image{
		src: rgb,
	}
}
