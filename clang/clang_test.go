package clang_test

import (
	//"fmt"
	"github.com/blinkat/ui/clang"
	"testing"
)

type TestWin struct {
	*clang.Module
	b bool
}

func (t *TestWin) OnPaint(c *clang.Canvas) {
	t.Module.OnPaint(c)

	c.FillStyle = clang.DefaultLineColor
	//c.FillRoundRect(clang.NewRect(100, 100, 300, 300), 12)

	c.DrawText("我操123", clang.NewPoint(20, 20), 12)

	// test draw img
	img, _ := clang.LoadImage("../test/2.png")
	c.DrawImage(img, clang.NewPoint(100, 500))

	// draw path
	c.BeginPath(clang.NewPoint(0, 600))
	c.Beizer2(clang.NewPoint(100, 600), clang.NewPoint(120, 500), clang.NewPoint(300, 600))
	c.LineMoveTo(clang.NewPoint(100, 700))
	c.Beizer3(clang.NewPoint(100, 700), clang.NewPoint(120, 700), clang.NewPoint(120, 720), clang.NewPoint(300, 700))

	c.LineMoveTo(clang.NewPoint(100, 100))
	c.Rect(clang.NewRect(100, 100, 300, 300))
	c.LineMoveTo(clang.NewPoint(500, 100))
	c.RoundRect(clang.NewRect(500, 100, 300, 300), 12)

	c.EndPath()
	c.Line.Color = clang.NewRGB(200, 80, 80)
	c.Line.Width = 1
	c.Stroke()
}

func TestHover(t *testing.T) {
	clang.Init()
	m, _ := clang.NewModule(clang.NewSize(1280, 800), clang.NewPoint(100, 100), 0, nil)
	win := &TestWin{
		Module: m,
		b:      false,
	}
	ico, _ := clang.LoadImage("../test/test.png")
	win.SetIcon(ico)

	clang.RunLoop(win)
}
