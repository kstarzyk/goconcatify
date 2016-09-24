package concatify

import (
	"fmt"
	"image"
	"os"
)

func openAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer imgFile.Close()

	img, format, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	return img, format, nil
}

func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}

	}

	return pixels
}
