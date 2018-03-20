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

func readImagesFromPaths(paths []string) ([]image.Image, error) {
	var images = make([]image.Image, len(paths))
	for i, path := range paths {
		img, _, err := openAndDecode(path)
		if err != nil {
			return nil, err
		}
		if img != nil {
			images[i] = img
		}
	}

	return images, nil
}

func decodePixelsFromImage(img image.Image, offsetX, offsetY int) (pixels []*pixel) {
	dim := img.Bounds().Max
	pixels = make([]*pixel, dim.Y*dim.X)
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			pixels[y*dim.Y+x] = &pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
		}

	}

	return pixels
}
