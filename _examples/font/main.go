package main

import (
	"io/ioutil"
	"log"

	"github.com/eugenezinoviev/gopie"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	chart := &gopie.PieChart{
		Width:  500,
		Height: 300,
		Values: []gopie.Value{
			gopie.Value{Value: 15, Label: "ERROR"},
			gopie.Value{Value: 67, Label: "WARNING"},
			gopie.Value{Value: 47, Label: "SUGGESTION"},
			gopie.Value{Value: 609, Label: "HINT"},
		},
		SliceColors:    []string{"firebrick", "sandybrown", "mediumseagreen", "deepskyblue"},
		StrokeWidth:    2,
		LabelLineWidth: 1,
		EmbedFont:      true,
		FontFamily:     "Roboto Regular",
		FontSize:       14,
	}

	robotoRegular, err := ioutil.ReadFile("Roboto-Regular.ttf")
	check(err)

	chart.SetFont(robotoRegular)

	svgBytes, err := chart.SVG()
	check(err)

	ioutil.WriteFile("./chart.svg", svgBytes, 0666)
}
