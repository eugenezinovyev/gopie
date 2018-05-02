package gopie

import (
	"bytes"
	"encoding/base64"
	"text/template"

	"github.com/satori/go.uuid"
)

// SVG function renders pie chart to SVG
func (o PieChart) SVG() (svgBytes []byte, err error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return
	}

	longestLabelRect, err := measureLongestLabel(o)
	if err != nil {
		return
	}

	pieRect := calculatePieRect(o, longestLabelRect)
	labels := createLabels(o, pieRect)

	font, err := createFontDefault(o)
	if err != nil {
		return
	}

	c := chart{
		Labels:       labels,
		Width:        o.getWidth(),
		Height:       o.getHeight(),
		NeedsMasking: o.getStrokeWidth() > 0,
		EmbedFont:    o.EmbedFont,
		Font:         font,
	}

	if o.getInnerRadius() > 0 {
		donut := newDonut(o, pieRect, uid.String())
		c.Donut = &donut
	} else {
		pie := newPie(o, pieRect, uid.String())
		c.Pie = &pie
	}

	tpl, err := createSvgTemplate()

	if err != nil {
		return
	}

	svgBytes, err = renderTemplate(tpl, c)

	return
}

func createFontDefault(chart PieChart) (f fontDetails, err error) {
	var buffer bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)
	defer encoder.Close()
	_, err = encoder.Write(chart.getFontBytes())
	if err != nil {
		return
	}

	f = fontDetails{
		FontFamily: chart.getFontFamily(),
		Base64:     buffer.String(),
	}
	return
}

func renderTemplate(t *template.Template, c chart) (svgBytes []byte, err error) {
	var buffer bytes.Buffer
	err = t.Execute(&buffer, c)

	if err != nil {
		return
	}

	svgBytes = buffer.Bytes()
	return
}

func calculatePieRect(chart PieChart, longestLabelRect rectangle) rectangle {
	maxRect := calculatePieMaxRect(chart, longestLabelRect)

	if maxRect.Height < maxRect.Width {
		widthDiff := maxRect.Width - maxRect.Height
		return rectangle{
			Left:   maxRect.Left + widthDiff/2,
			Top:    maxRect.Top,
			Width:  maxRect.Width - widthDiff,
			Height: maxRect.Height,
		}
	}
	heightDiff := maxRect.Height - maxRect.Width
	return rectangle{
		Left:   maxRect.Left,
		Top:    maxRect.Top + heightDiff/2,
		Width:  maxRect.Width,
		Height: maxRect.Height - heightDiff,
	}
}

func calculatePieMaxRect(chart PieChart, longestLabelRect rectangle) rectangle {
	labelLineOuterLength := chart.getLabelLine()
	xOffset := longestLabelRect.Width + labelLineOuterLength
	yOffset := longestLabelRect.Height + labelLineOuterLength

	return rectangle{
		Left:   xOffset,
		Top:    yOffset,
		Width:  chart.getWidth() - xOffset*2,
		Height: chart.getHeight() - yOffset*2,
	}
}
