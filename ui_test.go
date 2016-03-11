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

	win, err := ui.NewWindow(1024, 768, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.RunMainLoop(win)
}
