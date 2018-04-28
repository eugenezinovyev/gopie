package gopie

import (
	"fmt"
	"math"
)

type pie struct {
	Slices     *[]slice
	Circle     *circle
	Background *circle
}

func newPie(chart PieChart, pieRect rect) pie {
	if len(chart.Values) == 0 {
		return emptyPie(chart, pieRect)
	}

	if len(chart.Values) == 1 {
		return uncutPie(chart, pieRect)
	}

	return cutPie(chart, pieRect)
}

func emptyPie(chart PieChart, pieRect rect) pie {
	return pie{}
}

func uncutPie(chart PieChart, pieRect rect) pie {
	centerX, centerY := pieRect.getCenter()
	slicesCircleRadius := calculateSlicesRadius(chart, pieRect)

	return pie{
		Circle: &circle{
			CenterX: centerX,
			CenterY: centerY,
			Radius:  slicesCircleRadius,
			Style:   createSliceStyle(chart, 0),
		},
		Background: createBackgroundCircle(chart, pieRect),
	}
}

func cutPie(chart PieChart, pieRect rect) pie {
	centerX, centerY := pieRect.getCenter()
	radius := calculateSlicesRadius(chart, pieRect)

	sum := float64(0)
	total := chart.calculateTotalValue()
	slices := make([]slice, len(chart.Values))
	for index, value := range chart.Values {
		angleOffset := (sum / total) * twoPi
		slices[index] = createSlice(index, total, angleOffset, value, centerX, centerY, radius, chart)
		sum += value.Value
	}

	return pie{
		Slices:     &slices,
		Background: createBackgroundCircle(chart, pieRect),
	}
}

func createSlice(id int, total, angleOffset float64, value Value, centerX, centerY, radius float64, chart PieChart) slice {
	angle := twoPi * value.Value / total
	startX, startY := toDecartTranslate(angleOffset, radius, centerX, centerY)
	endX, endY := toDecartTranslate(angleOffset+angle, radius, centerX, centerY)

	lineSweep := 0
	if angle > math.Pi {
		lineSweep = 1
	}

	path := fmt.Sprintf(
		"M %v %v L %v %v A %v %v 0 %v 1 %v %v Z",
		centerX, centerY,
		startX, startY,
		radius, radius,
		lineSweep,
		endX, endY)

	return slice{
		ID:    id,
		Path:  path,
		Style: createSliceStyle(chart, id),
	}
}

func createSliceStyle(chart PieChart, id int) style {
	return style{
		Fill:        chart.getSliceColor(id),
		StrokeColor: chart.getStrokeColor(),
		StrokeWidth: chart.getStrokeWidth(),
	}
}

func createBackgroundCircle(chart PieChart, pieRect rect) *circle {
	centerX, centerY := pieRect.getCenter()
	radius := pieRect.calculateIncircleRadius()

	return &circle{
		CenterX: centerX,
		CenterY: centerY,
		Radius:  radius,
		Style: style{
			Fill: chart.getStrokeColor(),
		},
	}
}

func calculateSlicesRadius(c PieChart, r rect) float64 {
	outerRadius := r.calculateIncircleRadius()
	strokeWidth := c.getStrokeWidth()
	if strokeWidth == 0 {
		return outerRadius
	}
	return outerRadius - float64(strokeWidth)/2
}
