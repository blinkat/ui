package ui

import (
	"fmt"
)

func init() {
	//modules = make()
}

var (
	DefaultBackground = NewSolidBrush(23, 23, 20, 1)
	DefaultDark       = NewSolidBrush(22, 23, 19, 1)
	DefaultBorder     = NewSolidBrush(81, 83, 74, 1)
)

const (
	_LOG_INFO = iota
	_LOG_ERR

	_NAME_SPACE = "blinkat/ui"
)

func putError(s string) error {
	e := fmt.Errorf("%s error: %s", _NAME_SPACE, s)
	log(e.Error(), _LOG_ERR)
	return e
}

func log(s string, t int) {
	var fix string
	switch t {
	case _LOG_INFO:
		fix = "[Info]"
	case _LOG_ERR:
		fix = "[Error]"
	}

	fmt.Println(_NAME_SPACE, fix, ":", s)
}

func info(s ...interface{}) {
	log(fmt.Sprint(s...), _LOG_INFO)
}
