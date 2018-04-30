# gopie
Package gopie renders pie charts to SVG.
# Installation
To install gopie package run:
> go get -u github.com/eugenezinoviev/gopie
# Dependencies
This package depends on the following packages:
- github.com/golang/freetype/truetype
- golang.org/x/image/font

These package are used for chart label dimensions measure. Check limitations of [github.com/golang/freetype](https://github.com/golang/freetype) package if you plan to use label font different than the default one.
# Defaults
- Font used for label size measurement: *Roboto Medium*
- Default slice colors: *Material Design palette 400* (see: [https://material.io/color](https://material.io/color))
- Stroke color: *white*
- Label font family: *Roboto Medium*
- Label font size: *12*
- Label line length: *10px*
- Label line width: *2px*
- Label text padding: *4px*
- DPI ued for label size measurement: *92*
- Chart width: *200px*
- Chart height: *200px*
- Font: *Roboto Medium*
- Embed font: *false*
# Usage
``` go
	chart := gopie.PieChart{
		Values: []gopie.Value{
			gopie.Value{Value: 1, Label: "One"},
			gopie.Value{Value: 2, Label: "Two"},
			gopie.Value{Value: 5, Label: "Five"},
		},
	}

	svgBytes, err := chart.SVG()
```
## Custom slice colors
Check example in `./_examples/custom_colors`
``` go
	chart := gopie.PieChart{
		...
		SliceColors:    []string{"firebrick", "sandybrown", "mediumseagreen", "deepskyblue"},
		...
	}
```
## Embed font
Check example in `./_examples/font`
``` go
	robotoRegular, err := ioutil.ReadFile("Roboto-Regular.ttf")
	check(err)
	chart.SetFont(robotoRegular)
```
You can also check `./_examples` folder for more examples of usage.
# Licence
MIT