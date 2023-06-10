package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}
	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()

	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	
	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	// got := rectangle.Area()
	// 	want := 100.0

	// 	// if got != want {
	// 	// 	t.Errorf("got %.2f want %.2f", got, want)
	// 	// }

	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	// got := circle.Area()
	// 	want := 314.1592653589793

	// 	// if got != want {
	// 	// 	t.Errorf("got %g want %g", got, want)
	// 	// }	

	// 	checkArea(t, circle, want)
	// })

	// Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{width: 12, height: 6}, want: 72.0},
        {shape: Circle{radius: 10}, want: 314.1592653589793},
        {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
