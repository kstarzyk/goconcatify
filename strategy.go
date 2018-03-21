package concatify

import "image"

type concatStrategy func([]image.Image, concatParams) (pixels []*pixel, w int, h int)

func verticalConcatStrategy(images []image.Image, params concatParams) (pixels []*pixel, w int, h int) {
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

func horizontalConcatStrategy(images []image.Image, params concatParams) (pixels []*pixel, w int, h int) {
	w, h = 0, 0
	_, offsetH := params.GetOffset()
	for _, img := range images {
		imgPixels := decodePixelsFromImage(img, w, offsetH)
		if img.Bounds().Max.Y > h {
			h = img.Bounds().Max.Y
		}
		w += img.Bounds().Max.X

		pixels = append(pixels, imgPixels...)
	}

	return pixels, w, h
}

func gridConcatStrategy(images []image.Image, params concatParams) (pixels []*pixel, w int, h int) {
	w, h = 0, 0
	cols, rows := int(params.Cols), int(params.Rows)
	for i := 0; i < rows; i++ {
		slice := images[i*cols : i*cols+cols]
		params.SetOffset(0, h)
		rowPixels, rowW, rowH := horizontalConcatStrategy(slice, params)
		pixels = append(pixels, rowPixels...)

		h += rowH
		if rowW > w {
			w = rowW
		}
	}

	return pixels, w, h
}
