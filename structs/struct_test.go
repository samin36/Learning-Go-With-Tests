package structsAndinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g but wanted %g", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{6.0, 12.0}, 72.0},
		{Circle{3.0}, math.Pi * math.Pow(3.0, 2)},
		{Triangle{2.0, 4.0}, 4.0},
	}

	for _, aT := range areaTests {
		got := aT.shape.Area()
		if got != aT.want {
			t.Errorf("got %g but wanted %g for shape %#v", got, aT.want, aT.shape)
		}
	}
}
