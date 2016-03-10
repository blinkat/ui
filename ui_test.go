package ui_test

import (
	"fmt"
	"github.com/blinkat/ui"
	"testing"
)

func TestBase(t *testing.T) {
	err := ui.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	img, err := ui.NewImageForFile("./test/test.png")
	w, err := ui.CreateCustomWindow(1280, 768, 100, 100, "测试窗口", img, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Show()
}
