package concatify

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Arrangment string

var orientation Arrangment

const (
	Vertical   Arrangment = "vertical"
	Horizontal Arrangment = "horizontal"
	Grid       Arrangment = "grid"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

type ConcatImage struct {
	Sources    []string
	Strategy   ConcatStrategy
	Params     ConcatParams
	finalImage *image.RGBA
}

type ConcatParams struct {
	Display  Arrangment
	SameSize bool
	Rows     int
	Cols     int
}

type ConcatStrategy func([]image.Image, ConcatParams) ([]Pixel, int, int)

func verticalConcatStrategy(images []image.Image, params ConcatParams) (pixels []Pixel, w int, h int) {
	w, h = 0, 0
	for _, img := range images {
		if img == nil {
			continue
		}
		imgPixels := decodePixelsFromImage(img, 0, h)
		if img.Bounds().Max.X > w {
			w = img.Bounds().Max.X
		}
		h += img.Bounds().Max.Y
		pixels = append(pixels, imgPixels...)
	}
	return pixels, w, h
}

func horizontalConcatStrategy(images []image.Image, params ConcatParams) (pixels []Pixel, w int, h int) {
	w, h = 0, 0
	for _, img := range images {
		imgPixels := decodePixelsFromImage(img, w, 0)
		if img.Bounds().Max.Y > h {
			h = img.Bounds().Max.Y
		}
		w += img.Bounds().Max.X
		pixels = append(pixels, imgPixels...)
	}

	return pixels, w, h
}

func gridConcatStrategy(images []image.Image, params ConcatParams) (pixels []Pixel, w int, h int) {
	w, h = 0, 0
	for _, img := range images {
		imgPixels := decodePixelsFromImage(img, 0, h)
		if img.Bounds().Max.X > w {
			w = img.Bounds().Max.X
		}
		h += img.Bounds().Max.Y
		pixels = append(pixels, imgPixels...)
	}
	return pixels, w, h
}

func NewConcatImage(sources []string, params ConcatParams) (*ConcatImage, error) {

	cimg := &ConcatImage{}
	cimg.Sources = sources
	cimg.Params = params

	switch params.Display {
	case Vertical:
		cimg.Strategy = verticalConcatStrategy
	case Horizontal:
		cimg.Strategy = horizontalConcatStrategy
	case Grid:
		cimg.Strategy = gridConcatStrategy
	default:
		panic(params.Display)
	}

	err := cimg.draw()

	return cimg, err
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
