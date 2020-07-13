package main

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height +rectangle.Width)
}

//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}

func (c Circle)Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64  {
    return r.Width * r.Height
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle)Area() float64 {
	return (t.Base * t.Height) * 0.5
}

