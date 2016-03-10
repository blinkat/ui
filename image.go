package ui

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Image struct {
	Pix    []uint8
	Width  int
	Height int
	Icon   Icon
}

func (i *Image) Destory() {
	cDestoryIcon(i.Icon)
}

func NewImage(img image.Image) (*Image, error) {
	ret := &Image{}

	switch img.(type) {
	case *image.NRGBA:
		i := img.(*image.NRGBA)
		ret.Pix = i.Pix
	case *image.RGBA:
		i := img.(*image.RGBA)
		ret.Pix = i.Pix
	case *image.NRGBA64:
		i := img.(*image.NRGBA64)
		ret.Pix = i.Pix
	case *image.RGBA64:
		i := img.(*image.RGBA64)
		ret.Pix = i.Pix

	default:
		return nil, fmt.Errorf("image muse be image.NRGBA or image.RGBA.")
	}

	ret.Width = img.Bounds().Size().X
	ret.Height = img.Bounds().Size().Y

	ico := cLoadIcon(ret.Pix, ret.Width, ret.Height)
	if ico == nil {
		return nil, fmt.Errorf("failed load image.")
	}
	ret.Icon = ico
	return ret, nil
}

func NewImageForFile(path string) (*Image, error) {
	f, err := os.Open("./test/test.png")
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return NewImage(img)
}
