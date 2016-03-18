package clang

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Size struct {
	Width, Height int
}

func NewSize(w, h int) *Size {
	return &Size{
		Width:  w,
		Height: h,
	}
}

// *************** rectangle ****************
type Rectangle struct {
	X, Y, Width, Height int
}

func (r *Rectangle) Left() int {
	return r.X
}

func (r *Rectangle) Right() int {
	return r.X + r.Width
}

func (r *Rectangle) Top() int {
	return r.Y
}

func (r *Rectangle) Bottom() int {
	return r.Y + r.Height
}

// point int rect
func (r *Rectangle) PointIn(p *Point) bool {
	return (p.X >= r.Left() && p.Y >= r.Top()) && (p.X < r.Right() && p.Y < r.Bottom())
}

func Rect(left, top, right, bottom int) *Rectangle {
	return &Rectangle{
		X:      left,
		Y:      top,
		Width:  right - left,
		Height: bottom - top,
	}
}

func NewRect(x, y, width, height int) *Rectangle {
	return &Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
