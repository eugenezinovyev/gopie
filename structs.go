package gopie

import "math"

type style struct {
	Fill        string
	StrokeWidth float64
	StrokeColor string
}

type slice struct {
	ID    int
	Path  string
	Style style
}

type circle struct {
	CenterX float64
	CenterY float64
	Radius  float64
	Style   style
}

type line struct {
	X1    float64
	Y1    float64
	X2    float64
	Y2    float64
	Style style
}

type text struct {
	Text       string
	X          float64
	Y          float64
	FontSize   float64
	FontFamily string
	FontAnchor string
}

type chart struct {
	Width        float64
	Height       float64
	Pie          *pie
	Donut        *donut
	Labels       []label
	NeedsMasking bool
	EmbedFont    bool
	Font         fontDetails
}

type fontDetails struct {
	FontFamily string
	Base64     string
}

type label struct {
	Text text
	Line line
}

type rectangle struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}

func (r rectangle) getCenter() (centerX, centerY float64) {
	return r.Left + r.Width/2, r.Top + r.Height/2
}

func (r rectangle) calculateIncircleRadius() float64 {
	return math.Min(r.Height, r.Width) / 2
}
