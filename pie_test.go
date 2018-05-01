package gopie

import "testing"

func TestCalculateSlicesRadiusNoStroke(t *testing.T) {
	r := rectangle{Left: 0, Top: 0, Width: 100, Height: 100}
	chart := PieChart{StrokeWidth: 0}

	radius := calculateSlicesRadius(chart, r)
	expected := r.Width / 2

	if radius != expected {
		t.Fatalf("Expected %v but found %v", expected, radius)
	}
}

func TestCalculateSlicesRadiusWithStroke(t *testing.T) {
	strokeWidth := 10.0
	r := rectangle{Left: 0, Top: 0, Width: 100, Height: 100}
	chart := PieChart{StrokeWidth: strokeWidth}

	radius := calculateSlicesRadius(chart, r)
	expected := r.Width/2 - strokeWidth/2

	if radius != expected {
		t.Fatalf("Expected %v but found %v", expected, radius)
	}
}

func TestCreateBackgroundCircle(t *testing.T) {
	chart := PieChart{}
	rect := rectangle{Left: 0, Top: 0, Width: 100, Height: 100}

	actualCircle := createBackgroundCircle(chart, rect)
	expectedCircle := circle{
		CenterX: 50,
		CenterY: 50,
		Radius:  50,
		Style: style{
			Fill: defaultStrokeColor,
		},
	}

	if *actualCircle != expectedCircle {
		t.Fatalf("Expected %v but found %v", *actualCircle, expectedCircle)
	}
}
