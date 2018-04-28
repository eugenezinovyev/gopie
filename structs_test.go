package gopie

import "testing"

func TestRectGetCenterAtZeroPoint(t *testing.T) {
	r := rect{Left: 0, Top: 0, Width: 10, Height: 15}

	cx, cy := r.getCenter()

	expectedX := float64(5)
	expectedY := float64(7.5)
	if !(cx == expectedX && cy == expectedY) {
		t.Fatalf("Expected rect center to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, cx, cy)
	}
}

func TestRectGetCenterShiftedFromZeroPoint(t *testing.T) {
	r := rect{Left: 14, Top: 16, Width: 10, Height: 15}

	cx, cy := r.getCenter()

	expectedX := float64(19)
	expectedY := float64(23.5)
	if !(cx == expectedX && cy == expectedY) {
		t.Fatalf("Expected rect center to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, cx, cy)
	}
}

func TestRectGetCenterShiftedToNegitive(t *testing.T) {
	r := rect{Left: -14, Top: -16, Width: 10, Height: 15}

	cx, cy := r.getCenter()

	expectedX := float64(-9)
	expectedY := float64(-8.5)
	if !(cx == expectedX && cy == expectedY) {
		t.Fatalf("Expected rect center to be on (%v, %v) but found (%v, %v)", expectedX, expectedY, cx, cy)
	}
}

func TestCalculateIncircleRadiusEqualDimesions(t *testing.T) {
	r := rect{Left: 0, Top: 0, Width: 10, Height: 10}

	expectedRadius := float64(5)
	radius := r.calculateIncircleRadius()

	if !(radius == expectedRadius) {
		t.Fatalf("Expected incirle radius be %v but found %v", expectedRadius, radius)
	}
}

func TestCalculateIncircleRadiusHeightGreater(t *testing.T) {
	r := rect{Left: 0, Top: 0, Width: 10, Height: 15}

	expectedRadius := float64(5)
	radius := r.calculateIncircleRadius()

	if !(radius == expectedRadius) {
		t.Fatalf("Expected incirle radius be %v but found %v", expectedRadius, radius)
	}
}

func TestCalculateIncircleRadiusWidthGreater(t *testing.T) {
	r := rect{Left: 0, Top: 0, Width: 15, Height: 10}

	expectedRadius := float64(5)
	radius := r.calculateIncircleRadius()

	if !(radius == expectedRadius) {
		t.Fatalf("Expected incirle radius be %v but found %v", expectedRadius, radius)
	}
}
