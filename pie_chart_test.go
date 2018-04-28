package gopie

import (
	"reflect"
	"testing"
)

var (
	defaultPieChart = PieChart{}
	filledPieChart  = PieChart{
		Width:          112,
		Height:         113,
		DPI:            120,
		FontFamily:     "FFF",
		FontSize:       114,
		LabelLine:      111,
		LabelLineWidth: 110,
		LabelPadding:   109,
		SliceColors:    []string{"color1", "color2", "color3", "color4"},
		StrokeColor:    "color5",
		StrokeWidth:    115,
		Values: []Value{
			Value{116, "Label1"},
			Value{117, "Label2"},
			Value{118, "Label3"},
			Value{119, "Label4"},
		},
	}
)

func TestGetFontFamilyDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultFontFamily
	actual := chart.getFontFamily()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetFontSizeDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultFontSize
	actual := chart.getFontSize()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineFullLengthDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultLabelLine
	actual := chart.getLabelLineFullLength()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultLabelLine
	actual := chart.getLabelLine()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineWidthDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultLabelLineWidth
	actual := chart.getLabelLineWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelPaddingDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultLabelPadding
	actual := chart.getLabelPadding()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetDPIDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultDPI
	actual := chart.getDPI()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetSliceColorsDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultColors
	actual := chart.getSliceColors()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Expected a pie chart to return default colors if not set")
	}
}

func TestGetSliceColorDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultColors[0]
	actual := chart.getSliceColor(0)

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetStrokeWidthDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultStrokeWidth
	actual := chart.getStrokeWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetStrokeColorDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultStrokeColor
	actual := chart.getStrokeColor()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetWidthDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultWidth
	actual := chart.getWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetHeightDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := defaultHeight
	actual := chart.getHeight()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetCenterDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expectedX := defaultWidth / 2
	expectedY := defaultHeight / 2
	actualX, actualY := chart.getCenter()

	if !(actualX == expectedX && actualY == expectedY) {
		t.Fatalf("Expected (%v, %v) but found (%v, %v)", expectedX, expectedY, actualX, actualY)
	}
}

func TestCalculateTotalValueDefaultValue(t *testing.T) {
	chart := defaultPieChart

	expected := float64(0)
	actual := chart.calculateTotalValue()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetFontFamilySpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.FontFamily
	actual := chart.getFontFamily()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetFontSizeSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.FontSize
	actual := chart.getFontSize()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineFullLengthSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.LabelLine + filledPieChart.StrokeWidth
	actual := chart.getLabelLineFullLength()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.LabelLine
	actual := chart.getLabelLine()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelLineWidthSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.LabelLineWidth
	actual := chart.getLabelLineWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetLabelPaddingSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.LabelPadding
	actual := chart.getLabelPadding()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetDPISpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.DPI
	actual := chart.getDPI()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetSliceColorsSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.SliceColors
	actual := chart.getSliceColors()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Expected a pie chart to return default colors if not set")
	}
}

func TestGetSliceColorSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.SliceColors[0]
	actual := chart.getSliceColor(0)

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetStrokeWidthSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.StrokeWidth
	actual := chart.getStrokeWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetStrokeColorSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.StrokeColor
	actual := chart.getStrokeColor()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetWidthSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.Width
	actual := chart.getWidth()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetHeightSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := filledPieChart.Height
	actual := chart.getHeight()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}

func TestGetCenterSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expectedX := filledPieChart.Width / 2
	expectedY := filledPieChart.Height / 2
	actualX, actualY := chart.getCenter()

	if !(actualX == expectedX && actualY == expectedY) {
		t.Fatalf("Expected (%v, %v) but found (%v, %v)", expectedX, expectedY, actualX, actualY)
	}
}

func TestCalculateTotalValueSpecifiedValue(t *testing.T) {
	chart := filledPieChart

	expected := float64(470)
	actual := chart.calculateTotalValue()

	if actual != expected {
		t.Fatalf("Expected %v but found %v", expected, actual)
	}
}
