/*
	Package concatify implements a high-level API to concat images.
*/
package concatify

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type arrangment string

var orientation arrangment

const (
	vertical   arrangment = "vertical"
	horizontal arrangment = "horizontal"
	grid       arrangment = "grid"
)

type pixel struct {
	Point image.Point
	Color color.Color
}

type concatParams struct {
	Display  arrangment
	SameSize bool
	Rows     int64
	Cols     int64
	offsetW  int
	offsetH  int
}

func defaultParams(display arrangment) concatParams {
	switch display {
	case grid:
		return concatParams{"grid", true, 1, 1, 0, 0}
	case horizontal:
		return concatParams{"horizontal", true, 1, 1, 0, 0}
	case vertical:
		return concatParams{"vertical", true, 1, 1, 0, 0}
	}
	return concatParams{"vertical", true, 1, 1, 0, 0}

}

func (cp *concatParams) SetOffset(offW, offH int) {
	cp.offsetW = offW
	cp.offsetH = offH
}

func (cp *concatParams) SetDimension(rows, cols int64) {
	cp.Rows = rows
	cp.Cols = cols
}

func (cp concatParams) GetOffset() (int, int) {
	return cp.offsetW, cp.offsetH
}

func new(sources []string, params concatParams) (*ConcatImage, error) {
	cimg := &ConcatImage{}
	cimg.Sources = sources
	cimg.Params = params

	switch params.Display {
	case vertical:
		cimg.Strategy = verticalConcatStrategy
	case horizontal:
		cimg.Strategy = horizontalConcatStrategy
	case grid:
		cimg.Strategy = gridConcatStrategy
	default:
		panic(params.Display)
	}

	err := cimg.draw()

	return cimg, err
}

// ConcatImageContain
type ConcatImage struct {
	Sources    []string
	Strategy   concatStrategy
	Params     concatParams
	finalImage *image.RGBA
}

func (cimg *ConcatImage) draw() error {
	images, err := readImagesFromPaths(cimg.Sources)
	if err != nil {
		return err
	}

	pixels, w, h := cimg.Strategy(images, cimg.Params)

	newRect := image.Rectangle{
		Min: images[0].Bounds().Min,
		Max: image.Point{
			X: w,
			Y: h,
		},
	}

	cimg.finalImage = image.NewRGBA(newRect)
	for _, px := range pixels {
		cimg.finalImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}
	draw.Draw(cimg.finalImage, cimg.finalImage.Bounds(), cimg.finalImage, image.Point{0, 0}, draw.Src)
	return nil
}

// Save file to path
func (cimg *ConcatImage) Save(path string) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, cimg.finalImage)
	if err != nil {
		panic(err)
	}
}

func NewGrid(sources []string, rows, cols int64) (*ConcatImage, error) {
	var params concatParams
	params = defaultParams("grid")
	if len(sources) != int(rows*cols) {
		return nil, fmt.Errorf("#sources = %d, %dx%d = %d", len(sources), rows, cols, rows*cols)
	}
	params.SetDimension(rows, cols)
	return new(sources, params)
}

func NewVertical(sources []string) (*ConcatImage, error) {
	return new(sources, defaultParams("vertical"))
}

func NewHorizontal(sources []string) (*ConcatImage, error) {
	return new(sources, defaultParams("horizontal"))
}
