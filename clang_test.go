package ui_test

import (
	//"fmt"
	"github.com/blinkat/ui"
	"testing"
)

type TestWin struct {
	*ui.Module
	b bool
}

func (t *TestWin) OnPaint(c *ui.Canvas) {
	t.Module.OnPaint(c)

	c.FillStyle = ui.DefaultLineColor
	//c.FillRoundRect(ui.NewRect(100, 100, 300, 300), 12)

	c.DrawText("我操123", ui.NewPoint(20, 20), 12)

	// test draw img
	img, _ := ui.LoadImage("./test/2.png")
	c.DrawImage(img, ui.NewPoint(100, 500))

	// draw path
	c.BeginPath(ui.NewPoint(0, 600))
	c.Beizer2(ui.NewPoint(100, 600), ui.NewPoint(120, 500), ui.NewPoint(300, 600))
	c.LineMoveTo(ui.NewPoint(100, 700))
	c.Beizer3(ui.NewPoint(100, 700), ui.NewPoint(120, 700), ui.NewPoint(120, 720), ui.NewPoint(300, 700))

	c.LineMoveTo(ui.NewPoint(100, 100))
	c.Rect(ui.NewRect(100, 100, 300, 300))
	c.LineMoveTo(ui.NewPoint(500, 100))
	c.RoundRect(ui.NewRect(500, 100, 300, 300), 12)

	c.EndPath()
	c.Line.Color = ui.NewRGB(200, 80, 80)
	c.Line.Width = 1
	c.Stroke()
}

func TestHover(t *testing.T) {
	ui.Init()
	m, _ := ui.NewModule(ui.NewSize(1280, 800), ui.NewPoint(100, 100), 0, nil)
	win := &TestWin{
		Module: m,
		b:      false,
	}
	ico, _ := ui.LoadImage("./test/test.png")
	win.SetIcon(ico)
	win.SetText("测试窗口. test--")

	ui.RunLoop(win)
}
