package gopie

import (
	"fmt"
	"math"
)

type donut struct {
	Slices      *[]slice
	OuterCircle *circle
	InnerCircle *circle
	Background  *circle
}

func newDonut(chart PieChart, donutRect rectangle, uid string) donut {
	if len(chart.Values) == 0 {
		return emptyDonut(chart, donutRect, uid)
	}

	if len(chart.Values) == 1 {
		return uncutDonut(chart, donutRect, uid)
	}

	return cutDonut(chart, donutRect, uid)
}

func emptyDonut(chart PieChart, donutRect rectangle, uid string) donut {
	return donut{}
}

func uncutDonut(chart PieChart, donutRect rectangle, uid string) donut {
	centerX, centerY := donutRect.getCenter()
	slicesCircleRadius := calculateSlicesRadius(chart, donutRect)

	return donut{
		OuterCircle: &circle{
			CenterX: centerX,
			CenterY: centerY,
			Radius:  slicesCircleRadius,
			Style:   createSliceStyle(chart, 0),
			ChartID: uid,
		},
		InnerCircle: &circle{
			CenterX: centerX,
			CenterY: centerY,
			Radius:  chart.getInnerRadius(),
			Style:   createDonutInnerCircleStyle(chart),
		},
		Background: createBackgroundCircle(chart, donutRect),
	}
}

func cutDonut(chart PieChart, donutRect rectangle, uid string) donut {
	centerX, centerY := donutRect.getCenter()
	outerRadius := calculateSlicesRadius(chart, donutRect)
	innerRadius := chart.getInnerRadius()

	sum := float64(0)
	total := chart.calculateTotalValue()
	slices := make([]slice, len(chart.Values))
	for index, value := range chart.Values {
		angleOffset := (sum / total) * twoPi
		slices[index] = createDonutSlice(uid, index, total, angleOffset, value, centerX, centerY, outerRadius, innerRadius, chart)
		sum += value.Value
	}

	return donut{
		Slices:     &slices,
		Background: createBackgroundCircle(chart, donutRect),
	}
}

func createDonutSlice(chartID string, id int, total, angleOffset float64, value Value, centerX, centerY, outerRadius, innerRadius float64, chart PieChart) slice {
	angle := twoPi * value.Value / total
	outerStartX, outerStartY := toDecartTranslate(angleOffset, outerRadius, centerX, centerY)
	outerEndX, outerEndY := toDecartTranslate(angleOffset+angle, outerRadius, centerX, centerY)
	innerStartX, innerStartY := toDecartTranslate(angleOffset, innerRadius, centerX, centerY)
	innerEndX, innerEndY := toDecartTranslate(angleOffset+angle, innerRadius, centerX, centerY)

	lineSweep := 0
	if angle > math.Pi {
		lineSweep = 1
	}

	path := fmt.Sprintf(
		"M %v %v A %v %v 0 %v 0 %v %v L %v %v A %v %v 0 %v 1 %v %v Z",
		innerEndX, innerEndY, // moveTo

		innerRadius, innerRadius, // innerArc
		lineSweep,
		innerStartX, innerStartY,

		outerStartX, outerStartY, // lineTo

		outerRadius, outerRadius, // outerArc
		lineSweep,
		outerEndX, outerEndY)

	return slice{
		ChartID: chartID,
		ID:      id,
		Path:    path,
		Style:   createSliceStyle(chart, id),
	}
}
