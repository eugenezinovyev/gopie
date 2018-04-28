package gopie

import (
	"math"
	"testing"
)

func TestCalculateLabelLineInnerRadiusStrokeNotZero(t *testing.T) {
	var stroke float64 = 10
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius - stroke

	c := PieChart{StrokeWidth: stroke}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateLabelLineInnerRadius(c, r)

	if radius != expected {
		t.Fatalf("Label line should start right on stroke. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateLabelLineInnerRadiusStrokeIsZero(t *testing.T) {
	var stroke float64 = 0
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius

	c := PieChart{StrokeWidth: stroke}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateLabelLineInnerRadius(c, r)

	if radius != expected {
		t.Fatalf("Label line should start on pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateLabelLineOuterRadiusStrokeNotZero(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 10
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius + labelLine

	c := PieChart{StrokeWidth: stroke, LabelLineWidth: labelLine}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateLabelLineOuterRadius(c, r)

	if radius != expected {
		t.Fatalf("Label line should end in <labelLineLength> away from pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateLabelLineOuterRadiusStrokeIsZero(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 10
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius + labelLine

	c := PieChart{StrokeWidth: stroke, LabelLineWidth: labelLine}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateLabelLineOuterRadius(c, r)

	if radius != expected {
		t.Fatalf("Label line should end in <labelLineLength> away from pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateTextRadiusNotZero(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 10
	var textPadding float64 = 10
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius + labelLine + textPadding

	c := PieChart{StrokeWidth: stroke, LabelLineWidth: labelLine, LabelPadding: textPadding}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateTextRadius(c, r)

	if radius != expected {
		t.Fatalf("Label text should be placed in <labelLineLength + textPadding> away from pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateTextRadiusIsZero(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 0
	var textPadding float64 = 10
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius + labelLine + textPadding

	c := PieChart{StrokeWidth: stroke, LabelLineWidth: labelLine, LabelPadding: textPadding}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateTextRadius(c, r)

	if radius != expected {
		t.Fatalf("Label text should be placed in <labelLineLength + textPadding> away from pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCalculateTextRadiusNoTextPadding(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 0
	var textPadding float64 = 0
	var rectIncircleRadius float64 = 50
	expected := rectIncircleRadius + labelLine + defaultLabelPadding

	c := PieChart{StrokeWidth: stroke, LabelLineWidth: labelLine, LabelPadding: textPadding}
	r := rect{Left: 0, Top: 0, Width: rectIncircleRadius * 2, Height: rectIncircleRadius * 2}
	radius := calculateTextRadius(c, r)

	if radius != expected {
		t.Fatalf("Label text should be placed in <labelLineLength + textPadding> away from pie border. Expected %v but found %v", expected, radius)
	}
}

func TestCreateLineDefaultsFallback(t *testing.T) {
	index := 1
	var innerRadius float64 = 10
	var outerRadius float64 = 10
	var angle float64 = 10
	expectedColor := defaultColors[index]
	expectedWidth := defaultLabelLineWidth

	chart := PieChart{}
	line := createLine(index, innerRadius, outerRadius, angle, chart)

	if line.Style.StrokeColor != expectedColor {
		t.Fatalf("Label line color should be selected from defaults when not specified. Expected \"%v\" but found \"%v\"", expectedColor, line.Style.StrokeColor)
	}

	if line.Style.StrokeWidth != expectedWidth {
		t.Fatalf("Label line width should be selected from defaults when not specified. Expected \"%v\" but found \"%v\"", expectedWidth, line.Style.StrokeWidth)
	}
}

func TestCreateLineStyled(t *testing.T) {
	index := 1
	var innerRadius float64 = 10
	var outerRadius float64 = 10
	var angle float64 = 10
	expectedColor := "expectedColor"
	expectedWidth := float64(12)

	chart := PieChart{
		SliceColors: []string{
			"color1",
			expectedColor,
		},
		LabelLineWidth: expectedWidth,
	}
	line := createLine(index, innerRadius, outerRadius, angle, chart)

	if line.Style.StrokeColor != expectedColor {
		t.Fatalf("Label line color should be selected from defaults when not specified. Expected \"%v\" but found \"%v\"", expectedColor, line.Style.StrokeColor)
	}

	if line.Style.StrokeWidth != expectedWidth {
		t.Fatalf("Label line width should be selected from defaults when not specified. Expected \"%v\" but found \"%v\"", expectedWidth, line.Style.StrokeWidth)
	}
}

func TestCreateLineCoordinatesCalculated(t *testing.T) {
	index := 1
	var innerRadius float64 = 10
	var outerRadius float64 = 10
	var angle float64 = 10
	var width float64 = 100
	var height float64 = 100
	centerX := width / 2
	centerY := height / 2
	expectedStartX := (math.Sin(angle) * innerRadius) + centerX
	expectedStartY := -(math.Cos(angle) * innerRadius) + centerY
	expectedEndX := (math.Sin(angle) * outerRadius) + centerX
	expectedEndY := -(math.Cos(angle) * outerRadius) + centerY

	chart := PieChart{Width: width, Height: height}
	line := createLine(index, innerRadius, outerRadius, angle, chart)

	if line.X1 != expectedStartX && line.Y1 != expectedStartY {
		t.Fatalf("Label line should start on (%v, %v) but found (%v, %v)", expectedStartX, expectedStartY, line.X1, line.Y1)
	}

	if line.X1 != expectedStartX && line.Y1 != expectedStartY {
		t.Fatalf("Label line should end on (%v, %v) but found (%v, %v)", expectedEndX, expectedEndY, line.X2, line.Y2)
	}
}

func TestCalculateLabelTextAnchorForAngle0(t *testing.T) {
	var angle float64 = 0
	expectedAnchor := "middle"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle45(t *testing.T) {
	var angle float64 = math.Pi / 4
	expectedAnchor := "start"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle90(t *testing.T) {
	var angle float64 = math.Pi / 2
	expectedAnchor := "start"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle135(t *testing.T) {
	var angle float64 = math.Pi / 4 * 3
	expectedAnchor := "start"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle180(t *testing.T) {
	var angle float64 = math.Pi
	expectedAnchor := "middle"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle225(t *testing.T) {
	var angle float64 = math.Pi / 4 * 5
	expectedAnchor := "end"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle270(t *testing.T) {
	var angle float64 = math.Pi / 2 * 3
	expectedAnchor := "end"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle315(t *testing.T) {
	var angle float64 = math.Pi / 4 * 7
	expectedAnchor := "end"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCalculateLabelTextAnchorForAngle360(t *testing.T) {
	var angle float64 = math.Pi * 2
	expectedAnchor := "middle"
	anchor := calculateLabelTextAnchor(angle)
	if anchor != expectedAnchor {
		t.Fatalf("Expected label anchor to be \"%v\" for angle %v but found \"%v\"", expectedAnchor, angle, anchor)
	}
}

func TestCreateTextCalculatesTextCoordinates(t *testing.T) {
	index := 0
	var textRadius float64 = 10
	var angle float64 = 10
	var width float64 = 100
	var height float64 = 100
	centerX := width / 2
	centerY := height / 2
	expectedX := (math.Sin(angle) * textRadius) + centerX
	expectedY := -(math.Cos(angle) * textRadius) + centerY

	chart := PieChart{Values: []Value{{1, "Label"}}, Width: width, Height: height}
	text := createText(index, textRadius, angle, chart)

	if text.X != expectedX && text.Y != expectedY {
		t.Fatalf("Label text should be placed on (%v, %v) but found (%v, %v)", expectedX, expectedY, text.X, text.Y)
	}
}

func TestCreateLabelsSingleLabel(t *testing.T) {
	var labelLine float64 = 10
	var stroke float64 = 10
	var width float64 = 100
	var height float64 = 100

	chart := PieChart{
		Width:       width,
		Height:      height,
		Values:      []Value{{1, "Label"}},
		LabelLine:   labelLine,
		StrokeWidth: stroke,
	}
	pieRect := rect{Left: 0, Top: 0, Width: 50, Height: 50}

	labels := createLabels(chart, pieRect)

	actualCount := len(labels)
	if actualCount != 1 {
		t.Fatalf("Expected only 1 label for a single value but found %v", actualCount)
	}

	actualLabel := labels[0]
	centerX := width / 2
	centerY := width / 2
	expectedX1 := centerX + pieRect.calculateIncircleRadius() - stroke
	expectedY1 := centerY
	expectedX2 := centerX + pieRect.calculateIncircleRadius() + labelLine
	expectedY2 := centerY
	if !(actualLabel.Line.X1 == expectedX1 && actualLabel.Line.X2 == expectedX2) {
		t.Fatalf("Expected (%v,%v)-(%v,%v) but found (%v,%v)-(%v,%v)",
			expectedX1, expectedY1,
			expectedX2, expectedY2,
			actualLabel.Line.X1, actualLabel.Line.Y1,
			actualLabel.Line.X2, actualLabel.Line.Y2)
	}
}

func TestCreateLabelsMultipleLabels(t *testing.T) {
	chart := PieChart{
		Values: []Value{{1, "Label1"}, {1, "Label2"}},
	}
	pieRect := rect{Left: 0, Top: 0, Width: 50, Height: 50}
	labels := createLabels(chart, pieRect)

	actualCount := len(labels)
	if actualCount != 2 {
		t.Fatalf("Expected 2 labels for two values but found %v", actualCount)
	}
}
