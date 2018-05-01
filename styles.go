package gopie

func createSliceStyle(chart PieChart, id int) style {
	return style{
		Fill:        chart.getSliceColor(id),
		StrokeColor: chart.getStrokeColor(),
		StrokeWidth: chart.getStrokeWidth(),
	}
}

func createBackgroundCircleStyle(chart PieChart) style {
	return style{
		Fill: chart.getStrokeColor(),
	}
}

func createDonutInnerCircleStyle(chart PieChart) style {
	return style{
		Fill: defaultBackgroundColor,
	}
}

func createLabelLineStyle(chart PieChart, id int) style {
	return style{
		StrokeWidth: chart.getLabelLineWidth(),
		StrokeColor: chart.getSliceColor(id),
	}
}
