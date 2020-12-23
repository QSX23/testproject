package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//GraphData just plots the data for the first 2 people in the set
func GraphData(x []Person) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Data"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Amount"

	err = plotutil.AddLinePoints(p,
		x[0].id, plotPoints(x[0]),
		x[1].id, plotPoints(x[1]),
		x[2].id, plotPoints(x[2]),
	)

	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 8*vg.Inch, "test.png"); err != nil {
		panic(err)
	}

}

//creates an x y plot of the data
func plotPoints(p Person) plotter.XYs {
	pts := make(plotter.XYs, len(p.value))

	for i := range pts {

		pts[i].X = float64(p.value[i].time)
		pts[i].Y = float64(p.value[i].num)
	}

	return pts

}
