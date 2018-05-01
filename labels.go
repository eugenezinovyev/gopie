package gopie

import (
	"math"
)

func createLabels(chart PieChart, pieRect rectangle) []label {
	if len(chart.Values) == 0 {
		return createEmptyLabels()
	}
	if len(chart.Values) == 1 {
		return createSingleLabel(chart, pieRect)
	}
	return createMultipleLabels(chart, pieRect)
}

func createEmptyLabels() []label {
	return make([]label, 0)
}

func createMultipleLabels(chart PieChart, pieRect rectangle) []label {
	labels := make([]label, len(chart.Values))

	sum := float64(0)
	total := chart.calculateTotalValue()
	labelLineInnerRadius := calculateLabelLineInnerRadius(chart, pieRect)
	labelLineOuterRadius := calculateLabelLineOuterRadius(chart, pieRect)
	textRadius := calculateTextRadius(chart, pieRect)

	for index, value := range chart.Values {
		labelLineAngle := ((sum + value.Value/2) / total) * twoPi

		labels[index] = label{
			Text: createText(index, textRadius, labelLineAngle, chart),
			Line: createLine(index, labelLineInnerRadius, labelLineOuterRadius, labelLineAngle, chart),
		}
		sum += value.Value
	}

	return labels
}

func createSingleLabel(chart PieChart, pieRect rectangle) []label {
	labelLineInnerRadius := calculateLabelLineInnerRadius(chart, pieRect)
	labelLineOuterRadius := calculateLabelLineOuterRadius(chart, pieRect)
	textRadius := calculateTextRadius(chart, pieRect)
	labelLineAngle := math.Pi / 2

	return []label{
		label{
			Text: createText(0, textRadius, labelLineAngle, chart),
			Line: createLine(0, labelLineInnerRadius, labelLineOuterRadius, labelLineAngle, chart),
		},
	}
}

func createText(index int, textRadius, angle float64, chart PieChart) text {
	centerX, centerY := chart.getCenter()
	textX, textY := toDecartTranslate(angle, textRadius, centerX, centerY)

	return text{
		Text:       chart.Values[index].Label,
		FontAnchor: calculateLabelTextAnchor(angle),
		FontFamily: chart.getFontFamily(),
		FontSize:   chart.getFontSize(),
		X:          textX,
		Y:          textY,
	}
}

func calculateLabelTextAnchor(angle float64) string {
	mod := math.Mod(angle, math.Pi*2)

	switch {
	case mod == 0:
		return "middle"
	case mod == math.Pi:
		return "middle"
	case mod > math.Pi:
		return "end"
	default:
		return "start"
	}
}

func createLine(index int, innerRadius, outerRadius, angle float64, chart PieChart) line {
	centerX, centerY := chart.getCenter()
	startX, startY := toDecartTranslate(angle, innerRadius, centerX, centerY)
	endX, endY := toDecartTranslate(angle, outerRadius, centerX, centerY)

	return line{
		Style: createLabelLineStyle(chart, index),
		X1:    startX,
		Y1:    startY,
		X2:    endX,
		Y2:    endY,
	}
}

func calculateLabelLineInnerRadius(c PieChart, r rectangle) float64 {
	outerRadius := r.calculateIncircleRadius()
	strokeWidth := c.getStrokeWidth()
	return outerRadius - strokeWidth
}

func calculateLabelLineOuterRadius(c PieChart, r rectangle) float64 {
	innerRadius := calculateLabelLineInnerRadius(c, r)
	return innerRadius + c.getLabelLineFullLength()
}

func calculateTextRadius(c PieChart, r rectangle) float64 {
	labelLineOuterRadius := calculateLabelLineOuterRadius(c, r)
	return labelLineOuterRadius + c.getLabelPadding()
}
