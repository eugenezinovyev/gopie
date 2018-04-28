package gopie

import (
	"math"
	"testing"
)

func TestPointsToPixels96DPI(t *testing.T) {
	dpi := float64(96)
	points := float64(3)

	pixels := pointsToPixels(dpi, points)
	expectedPixels := float64(4)

	if pixels != expectedPixels {
		t.Fatalf("Expected %vpt to be %vpx with %v DPI but found %vpx", points, expectedPixels, dpi, pixels)
	}
}

func TestPointsToPixels72DPI(t *testing.T) {
	dpi := float64(72)
	points := float64(3)

	pixels := pointsToPixels(dpi, points)
	expectedPixels := float64(3)

	if pixels != expectedPixels {
		t.Fatalf("Expected %vpt to be %vpx with %v DPI but found %vpx", points, expectedPixels, dpi, pixels)
	}
}

var eps = float64(0.00000001)

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func TestToDecartAngle0(t *testing.T) {
	angle := float64(0)
	radius := float64(10)
	expectedX := float64(0)
	expectedY := float64(-10)

	x, y := toDecart(angle, radius)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartAngle90(t *testing.T) {
	angle := math.Pi / 2
	radius := float64(10)
	expectedX := float64(10)
	expectedY := float64(0)

	x, y := toDecart(angle, radius)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartAngle45(t *testing.T) {
	angle := math.Pi / 4
	radius := float64(10)
	expectedX := 10 / math.Sqrt(2)
	expectedY := -(10 / math.Sqrt(2))

	x, y := toDecart(angle, radius)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartTranslateAngle0(t *testing.T) {
	dx := float64(13)
	dy := float64(11)
	angle := float64(0)
	radius := float64(10)
	expectedX := float64(0) + dx
	expectedY := float64(-10) + dy

	x, y := toDecartTranslate(angle, radius, dx, dy)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartTranslateTranslatesCoordinates(t *testing.T) {
	dx := float64(-5)
	dy := float64(-4)
	angle := float64(0)
	radius := float64(10)
	expectedX := float64(0) + dx
	expectedY := float64(-10) + dy

	x, y := toDecartTranslate(angle, radius, dx, dy)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartTranslateAngle90(t *testing.T) {
	dx := float64(13)
	dy := float64(11)
	angle := math.Pi / 2
	radius := float64(10)
	expectedX := float64(10) + dx
	expectedY := float64(0) + dy

	x, y := toDecartTranslate(angle, radius, dx, dy)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestToDecartTranslateAngle45(t *testing.T) {
	dx := float64(13)
	dy := float64(11)
	angle := math.Pi / 4
	radius := float64(10)
	expectedX := (10 / math.Sqrt(2)) + dx
	expectedY := (-(10 / math.Sqrt(2))) + dy

	x, y := toDecartTranslate(angle, radius, dx, dy)

	if !(floatEquals(x, expectedX) && floatEquals(y, expectedY)) {
		t.Fatalf("Expected point to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, x, y)
	}
}

func TestFindLongestLabelFindsLabelOnFirstPlace(t *testing.T) {
	longLabel := "theLongetLabel"
	chart := PieChart{
		Values: []Value{
			Value{Value: 1, Label: longLabel},
			Value{Value: 1, Label: "label1"},
			Value{Value: 1, Label: "label2"},
		},
	}

	label := findLongestLabel(chart)

	if label != longLabel {
		t.Fatalf("Expected the longetLabel to be %v but found %v", longLabel, label)
	}
}

func TestFindLongestLabelFindsLabelInMiddle(t *testing.T) {
	longLabel := "theLongetLabel"
	chart := PieChart{
		Values: []Value{
			Value{Value: 1, Label: "label1"},
			Value{Value: 1, Label: longLabel},
			Value{Value: 1, Label: "label2"},
		},
	}

	label := findLongestLabel(chart)

	if label != longLabel {
		t.Fatalf("Expected the longetLabel to be %v but found %v", longLabel, label)
	}
}

func TestFindLongestLabelFindsLabelOnLastPlace(t *testing.T) {
	longLabel := "theLongetLabel"
	chart := PieChart{
		Values: []Value{
			Value{Value: 1, Label: "label1"},
			Value{Value: 1, Label: longLabel},
			Value{Value: 1, Label: "label2"},
		},
	}

	label := findLongestLabel(chart)

	if label != longLabel {
		t.Fatalf("Expected the longetLabel to be %v but found %v", longLabel, label)
	}
}
