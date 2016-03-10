package ui

type Rect struct {
	Size
	Point
}

func (r *Rect) Left() int {
	return r.X
}

func (r *Rect) Top() int {
	return r.Y
}

func (r *Rect) Right() int {
	return r.X + r.Width
}

func (r *Rect) Bottom() int {
	return r.Y + r.Height
}

func NewRect(x, y, width, height int) Rect {
	return Rect{
		Point: NewPoint(x, y),
		Size:  NewSize(width, height),
	}
}

// *****************************
type Size struct {
	Width, Height int
}

func NewSize(width, height int) Size {
	return Size{
		Width:  width,
		Height: height,
	}
}

// ****************************
type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}
