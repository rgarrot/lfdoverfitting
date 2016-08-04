package main

import (
	"image/color"

	"golearn/num"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func plotBase(b Base, yP2 num.Matrix, yP10 num.Matrix) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plot Base"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(baseToPlotter(b))
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)

	// Make a scatter plotter and set its style.
	s, err = plotter.NewScatter(predictedToPlotter(b, yP2))
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 0, B: 128, A: 255}
	p.Add(s)

	// Make a scatter plotter and set its style.
	s, err = plotter.NewScatter(predictedToPlotter(b, yP10))
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 0, B: 0, G: 255, A: 255}
	p.Add(s)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func baseToPlotter(b Base) plotter.XYs {
	pts := make(plotter.XYs, len(b.Y))
	for i := range pts {
		pts[i].X = b.X[i]
		pts[i].Y = b.Y[i]
	}
	return pts
}

func predictedToPlotter(b Base, y num.Matrix) plotter.XYs {
	pts := make(plotter.XYs, len(b.X))
	for i := range pts {
		pts[i].X = b.X[i]
		pts[i].Y = y.Data[i][0]
	}
	return pts
}
