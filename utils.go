package concatify

import (
	"fmt"
	"image"
	"os"
)

func openAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	if err != nil {
		return nil, "", fmt.Errorf("Cannot open file %s", filepath)
	}
	defer imgFile.Close()

	img, format, err := image.Decode(imgFile)
	if err != nil {
		return nil, "", fmt.Errorf("Cannot decode file %s", filepath)
	}

	return img, format, nil
}

func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			pixels = append(pixels, &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			})
		}

	}

	return pixels
}
