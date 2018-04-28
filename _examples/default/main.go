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
		Values: []gopie.Value{
			gopie.Value{Value: 1, Label: "One"},
			gopie.Value{Value: 2, Label: "Two"},
			gopie.Value{Value: 5, Label: "Five"},
		},
	}

	svgBytes, err := chart.SVG()
	check(err)

	ioutil.WriteFile("./chart.svg", svgBytes, 0666)
}
