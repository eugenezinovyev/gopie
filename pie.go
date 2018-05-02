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

func newPie(chart PieChart, pieRect rectangle, uid string) pie {
	if len(chart.Values) == 0 {
		return emptyPie(chart, pieRect, uid)
	}

	if len(chart.Values) == 1 {
		return uncutPie(chart, pieRect, uid)
	}

	return cutPie(chart, pieRect, uid)
}

func emptyPie(chart PieChart, pieRect rectangle, uid string) pie {
	return pie{}
}

func uncutPie(chart PieChart, pieRect rectangle, uid string) pie {
	centerX, centerY := pieRect.getCenter()
	slicesCircleRadius := calculateSlicesRadius(chart, pieRect)

	return pie{
		Circle: &circle{
			CenterX: centerX,
			CenterY: centerY,
			Radius:  slicesCircleRadius,
			Style:   createSliceStyle(chart, 0),
			ChartID: uid,
		},
		Background: createBackgroundCircle(chart, pieRect),
	}
}

func cutPie(chart PieChart, pieRect rectangle, uid string) pie {
	centerX, centerY := pieRect.getCenter()
	radius := calculateSlicesRadius(chart, pieRect)

	sum := float64(0)
	total := chart.calculateTotalValue()
	slices := make([]slice, len(chart.Values))
	for index, value := range chart.Values {
		angleOffset := (sum / total) * twoPi
		slices[index] = createPieSlice(uid, index, total, angleOffset, value, centerX, centerY, radius, chart)
		sum += value.Value
	}

	return pie{
		Slices:     &slices,
		Background: createBackgroundCircle(chart, pieRect),
	}
}

func createPieSlice(chartID string, id int, total, angleOffset float64, value Value, centerX, centerY, radius float64, chart PieChart) slice {
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
		ChartID: chartID,
		ID:      id,
		Path:    path,
		Style:   createSliceStyle(chart, id),
	}
}

func createBackgroundCircle(chart PieChart, pieRect rectangle) *circle {
	centerX, centerY := pieRect.getCenter()
	radius := pieRect.calculateIncircleRadius()

	return &circle{
		CenterX: centerX,
		CenterY: centerY,
		Radius:  radius,
		Style:   createBackgroundCircleStyle(chart),
	}
}

func calculateSlicesRadius(c PieChart, r rectangle) float64 {
	outerRadius := r.calculateIncircleRadius()
	strokeWidth := c.getStrokeWidth()
	if strokeWidth == 0 {
		return outerRadius
	}
	return outerRadius - float64(strokeWidth)/2
}
