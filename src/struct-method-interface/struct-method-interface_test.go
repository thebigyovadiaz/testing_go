package struct_method_interface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape ShapePerimeter, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("Get rectangle perimeter", func(t *testing.T) {
		rectangle := RectangleAttr{
			Width:  10.0,
			Height: 10.0,
		}

		want := 40.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("Get circle perimeter", func(t *testing.T) {
		circle := CircleAttr{
			Radius: 10.0,
		}

		want := 62.83185307179586
		checkPerimeter(t, circle, want)
	})

	t.Run("Get triangle perimeter", func(t *testing.T) {
		triangle := TriangleAttr{
			Base:   6.0,
			Height: 10.0,
		}

		want := 27.6619037896906
		checkPerimeter(t, triangle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape ShapeArea, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Get rectangle area", func(t *testing.T) {
		rectangle := RectangleAttr{
			Width:  10.0,
			Height: 10.0,
		}

		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("Get circle area", func(t *testing.T) {
		circle := CircleAttr{
			Radius: 5.0,
		}

		want := 78.53981633974483
		checkArea(t, circle, want)
	})

	t.Run("Get triangle area", func(t *testing.T) {
		triangle := TriangleAttr{
			Base:   20.0,
			Height: 35.0,
		}

		want := 700.0
		checkArea(t, triangle, want)
	})
}

func TestAreaCases(t *testing.T) {
	areaTests := []struct {
		want  float64
		shape ShapeArea
	}{
		{100.0, RectangleAttr{10.0, 10.0}},
		{78.53981633974483, CircleAttr{5.0}},
		{700.0, TriangleAttr{20.0, 35.0}},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}
}

func TestPerimeterCases(t *testing.T) {
	parimeterTests := []struct {
		want  float64
		shape ShapePerimeter
	}{
		{40.0, RectangleAttr{10.0, 10.0}},
		{62.83185307179586, CircleAttr{10.0}},
		{27.6619037896906, TriangleAttr{6.0, 10.0}},
	}

	for _, tt := range parimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}
}
