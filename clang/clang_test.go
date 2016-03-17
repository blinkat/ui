package clang_test

import (
	//"fmt"
	"github.com/blinkat/ui/clang"
	"image"
	_ "image/png"
	"os"
	"testing"
)

type TestWin struct {
	*clang.Module
	b bool
}

func (t *TestWin) OnPaint(c *clang.Canvas) {
	t.Module.OnPaint(c)

	c.FillStyle = clang.DefaultLineColor
	c.FillRoundRect(100, 100, 300, 300, 12)

	c.DrawText("我操123", 20, 20, 12)

	f, err := os.Open("../test/2.png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	c.DrawImage(img, 100, 500)
}

func TestHover(t *testing.T) {
	clang.Init()
	m, _ := clang.NewModule(1280, 800, 100, 100, 0, nil)
	win := &TestWin{
		Module: m,
		b:      false,
	}

	clang.RunLoop(win)
}
