package gopie

import (
	"log"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func pointsToPixels(dpi, points float64) (pixels float64) {
	pixels = (points * dpi) / 72.0
	return
}

// toDecart converts polar coordinate to the Decard system with inverted Y-axis and start in left-top corner
func toDecart(angle, radius float64) (x, y float64) {
	return math.Sin(angle) * radius, -(math.Cos(angle) * radius)
}

// toDecartTranslate converts polar coordinate to the Decard system with inverted Y-axis and start in left-top corner
// and translated using dx and dy coordinates
func toDecartTranslate(angle, radius, dx, dy float64) (ex, ey float64) {
	sx, sy := toDecart(angle, radius)
	return sx + dx, sy + dy
}

func measureLongestLabel(chart PieChart) (rect, error) {
	longestLabel := findLongestLabel(chart)
	return measureString(chart, longestLabel, chart.getFontSize(), chart.getDPI())
}

func findLongestLabel(chart PieChart) string {
	longestLabel := ""
	for _, value := range chart.Values {
		if len(longestLabel) < len(value.Label) {
			longestLabel = value.Label
		}
	}
	return longestLabel
}

func measureString(chart PieChart, text string, fontSize, dpi float64) (r rect, err error) {
	f, err := chart.getFont()
	if err != nil {
		return
	}

	drawer := &font.Drawer{
		Face: truetype.NewFace(f, &truetype.Options{
			DPI:  dpi,
			Size: fontSize,
		}),
	}
	width := drawer.MeasureString(text).Ceil()
	heigth := float64(pointsToPixels(dpi, fontSize))

	r = rect{
		Width:  float64(width),
		Height: heigth,
	}
	return
}
