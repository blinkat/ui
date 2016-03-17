package clang

type Key int

func RunLoop(win Frame) {
	MainFrame = win
	win.Show()
	for win.GetMessage() {
	}

	for _, v := range hoverPool {
		v <- false
	}
}

var (
	DefaultBackground  = NewRGB(23, 24, 20)
	DefaultFrontground = NewRGB(39, 40, 34)
	DefaultLineColor   = NewRGB(60, 61, 56)
	DefaultFontColor   = NewRGB(255, 255, 255)
)
