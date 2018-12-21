package shape

import "github.com/go-gl/mathgl/mgl32"

type Rect struct {
	Mins mgl32.Vec2
	Maxs mgl32.Vec2
}

func (rect *Rect) X() float32 {
	return rect.Mins.X()
}

func (rect *Rect) Y() float32 {
	return rect.Mins.Y()
}

func (rect *Rect) Width() float32 {
	return rect.Maxs.X()
}

func (rect *Rect) Height() float32 {
	return rect.Maxs.Y()
}

func NewRect(mins mgl32.Vec2, maxs mgl32.Vec2) *Rect {
	return &Rect{
		Mins: mins,
		Maxs: maxs,
	}
}
