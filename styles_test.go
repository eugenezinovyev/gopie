package gopie

import "testing"

func TestCreateSliceStyleDefaultValue(t *testing.T) {
	id := 1
	chart := PieChart{}

	actualStyle := createSliceStyle(chart, id)
	expectedStyle := style{
		Fill:        defaultColors[id],
		StrokeColor: defaultStrokeColor,
		StrokeWidth: defaultStrokeWidth,
	}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}

func TestCreateSliceStyleSpecifiedValue(t *testing.T) {
	id := 1
	chart := PieChart{
		StrokeWidth: 13,
		StrokeColor: "StrokeColor",
		SliceColors: []string{"Slice0Color", "Slice1Color"},
	}

	actualStyle := createSliceStyle(chart, id)
	expectedStyle := style{
		Fill:        chart.SliceColors[id],
		StrokeColor: chart.StrokeColor,
		StrokeWidth: chart.StrokeWidth,
	}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}

func TestCreateBackgroundCircleStyleDefaultValue(t *testing.T) {
	chart := PieChart{}

	actualStyle := createBackgroundCircleStyle(chart)
	expectedStyle := style{Fill: defaultStrokeColor}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}

func TestCreateBackgroundCircleStyleSpecifiedValue(t *testing.T) {
	chart := PieChart{StrokeColor: "StrokeColor"}

	actualStyle := createBackgroundCircleStyle(chart)
	expectedStyle := style{Fill: chart.StrokeColor}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}

func TestCreateLabelLineStyleDefaultValue(t *testing.T) {
	id := 1
	chart := PieChart{}

	actualStyle := createLabelLineStyle(chart, id)
	expectedStyle := style{
		StrokeColor: defaultColors[id],
		StrokeWidth: defaultLabelLineWidth,
	}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}

func TestCreateLabelLineStyleSpecifiedValue(t *testing.T) {
	id := 1
	chart := PieChart{
		LabelLineWidth: 13,
		SliceColors:    []string{"Slice0Color", "Slice1Color"},
	}

	actualStyle := createLabelLineStyle(chart, id)
	expectedStyle := style{
		StrokeColor: chart.SliceColors[id],
		StrokeWidth: chart.LabelLineWidth,
	}

	if actualStyle != expectedStyle {
		t.Fatalf("Expected %v but found %v", expectedStyle, actualStyle)
	}
}
