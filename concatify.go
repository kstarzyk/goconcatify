package concatify

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

type direction uint8

const (
	VERTICAL direction = iota
	HORIZONTAL
)

type ConcatedImage struct {
	images    []image.Image
	pixels    []*Pixel
	format    string
	maxWidth  int
	maxHeight int
	ConcatedImageOptions
}

type ConcatedImageOptions struct {
	Direction  direction
	SameWidth  bool
	SameHeight bool
}

func getDefaultOptions() ConcatedImageOptions {
	return ConcatedImageOptions{
		Direction:  VERTICAL,
		SameWidth:  false,
		SameHeight: false,
	}
}
func NewConcatedImage(paths []string, options ...ConcatedImageOptions) *ConcatedImage {
	ci := &ConcatedImage{}
	var opt ConcatedImageOptions
	if len(options) == 0 {
		opt = getDefaultOptions()
	} else {
		opt = options[0]
	}
	ci.Direction = opt.Direction
	ci.maxHeight = 0
	ci.maxWidth = 0
	for _, path := range paths {
		img, _, err := openAndDecode(path)
		if err != nil {
			panic(err)
		}
		ci.images = append(ci.images, img)

		var imgPixels []*Pixel
		if ci.Direction == VERTICAL {
			imgPixels = DecodePixelsFromImage(img, 0, ci.maxHeight)
			if img.Bounds().Max.X > ci.maxWidth {
				ci.maxWidth = img.Bounds().Max.X
			}
			ci.maxHeight += img.Bounds().Max.Y
		} else {
			imgPixels = DecodePixelsFromImage(img, ci.maxWidth, 0)
			if img.Bounds().Max.Y > ci.maxHeight {
				ci.maxHeight = img.Bounds().Max.Y
			}
			ci.maxWidth += img.Bounds().Max.X
		}

		ci.pixels = append(ci.pixels, imgPixels...)

	}
	return ci
}

func (ci *ConcatedImage) Draw(dest string) {
	newRect := image.Rectangle{
		Min: ci.images[0].Bounds().Min,
		Max: image.Point{
			X: ci.maxWidth,
			Y: ci.maxHeight,
		},
	}

	finalImage := image.NewRGBA(newRect)

	for _, px := range ci.pixels {
		finalImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}
	draw.Draw(finalImage, finalImage.Bounds(), finalImage, image.Point{0, 0}, draw.Src)

	out, err := os.Create(dest)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	err = png.Encode(out, finalImage)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
