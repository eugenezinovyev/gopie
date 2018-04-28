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
	chart := gopie.PieChart{
		Width:  500,
		Height: 300,
		Values: []gopie.Value{
			gopie.Value{Value: 1, Label: "Single Value"},
		},
		LabelLineWidth: 1,
		FontSize:       14,
	}

	svgBytes, err := chart.SVG()
	check(err)

	ioutil.WriteFile("./chart.svg", svgBytes, 0666)
}
