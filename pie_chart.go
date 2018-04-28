/*
Package gopie renders pie charts to SVG.
*/
package gopie

//go:generate go run mbed/mbed.go -d ./assets -o ./assets/assets.go -p assets

// Value represents chart value.
type Value struct {
	Value float64 // Value.
	Label string  // Label of value.
}

// PieChart struct represents a pie chart rendering options.
type PieChart struct {
	Width          float64  // SVG Width. Default is 200.
	Height         float64  // SVG Height. Default is 200.
	Values         []Value  // Chart values.
	SliceColors    []string // Colors to be used for pie slices and label lines.
	StrokeColor    string   // Slice stroke color. Default is "#ffffff".
	StrokeWidth    float64  // Slice stroke width. Default is 0 px.
	FontSize       float64  // Label font size. Default is 12 px.
	FontFamily     string   // Label font family. Default is "Roboto Medium".
	LabelLine      float64  // Label line length. Default is 10 px.
	LabelLineWidth float64  // Label line width. Default is 2 px.
	LabelPadding   float64  // Label text padding. Default is 4 px.
	DPI            float64  // DPI. Default is 92.
}

func (c PieChart) getFontFamily() string {
	if c.FontFamily == "" {
		return defaultFontFamily
	}
	return c.FontFamily
}

func (c PieChart) getFontSize() float64 {
	if c.FontSize == 0 {
		return defaultFontSize
	}
	return c.FontSize
}

func (c PieChart) getLabelLineFullLength() float64 {
	stroke := c.getStrokeWidth()
	labelLine := c.getLabelLine()
	return labelLine + stroke
}

func (c PieChart) getLabelLine() float64 {
	if c.LabelLine == 0 {
		return defaultLabelLine
	}
	return c.LabelLine
}

func (c PieChart) getLabelLineWidth() float64 {
	if c.LabelLineWidth == 0 {
		return defaultLabelLineWidth
	}
	return c.LabelLineWidth
}

func (c PieChart) getLabelPadding() float64 {
	if c.LabelPadding == 0 {
		return defaultLabelPadding
	}
	return c.LabelPadding
}

func (c PieChart) getDPI() float64 {
	if c.DPI == 0 {
		return defaultDPI
	}
	return c.DPI
}

func (c PieChart) getSliceColors() []string {
	if len(c.SliceColors) == 0 {
		return defaultColors
	}
	return c.SliceColors
}

func (c PieChart) getSliceColor(index int) string {
	colors := c.getSliceColors()
	return colors[index%len(colors)]
}

func (c PieChart) getStrokeWidth() float64 {
	if c.StrokeWidth <= 0 {
		return defaultStrokeWidth
	}
	return c.StrokeWidth
}

func (c PieChart) getStrokeColor() string {
	if c.StrokeColor == "" {
		return defaultStrokeColor
	}
	return c.StrokeColor
}

func (c PieChart) getWidth() float64 {
	if c.Width == 0 {
		return defaultWidth
	}
	return c.Width
}

func (c PieChart) getHeight() float64 {
	if c.Height == 0 {
		return defaultHeight
	}
	return c.Height
}

func (c PieChart) getCenter() (centerX, centerY float64) {
	return c.getWidth() / 2, c.getHeight() / 2
}

func (c PieChart) calculateTotalValue() float64 {
	total := float64(0)

	for _, value := range c.Values {
		total = total + value.Value
	}

	return total
}
