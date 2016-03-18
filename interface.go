package ui

import (
	c "github.com/blinkat/ui/clang"
)

type Frame c.Frame
type Module c.Module

func NewModule(size *Size, p *Point, style int, parent Frame) (*Module, error) {
	return c.NewModule(size, p, style, parent)
}

// **********************
type Rect c.Rectangle
type Point c.Point
type Size c.Size
