package main

import (

	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.f want %.2f", got, want)
	}
}
type Shape interface {
    Area() float64
}

func TestArea(t *testing.T) {
 //「表格驱动测试」
	areaTeats := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10},314.1592653589793},
		{Triangle{12,6},36.0},
	}

	for _, tt := range areaTeats{
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
	//checkAreas := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{10.0, 10.0}
	//	checkAreas(t, rectangle, 100.0)
	//})
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkAreas(t,circle,314.1592653589793)
	//})
}
