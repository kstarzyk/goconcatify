package concatify_test

import (
	"fmt"
	"image"
	"os"
	"testing"

	concatify "github.com/kstarzyk/goconcatify"
)

func comparePNG(pathToFirst, pathToSecond string) bool {
	fst, fstFormat, err := openAndDecode(pathToFirst)
	if err != nil {
		return false
	}
	snd, sndFormat, err := openAndDecode(pathToSecond)
	if err != nil {
		return false
	}
	if fstFormat != sndFormat {
		return false
	}
	if fst.Bounds() != snd.Bounds() {
		return false
	}
	for y := 0; y < fst.Bounds().Max.Y; y++ {
		for x := 0; x < fst.Bounds().Max.X; x++ {
			if fst.At(x, y) != snd.At(x, y) {
				return false
			}
		}
	}
	return true
}

func remove(path string) {
	err := os.Remove(path)
	if err != nil {

	}
}

type concatifyTest struct {
	name          string
	paths         []string
	result, model string
}

type concatifyTestFail struct {
	concatifyTest
	expectedErrorMessage string
}

const (
	gopher           = "./mocks/gopher.png"
	aviator          = "./mocks/aviator.png"
	verticalResult   = "./mocks/golang-and-golang-aviator-ver.png"
	horizontalResult = "./mocks/golang-ang-golang-aviator-hor.png"
	gridResult       = "./mocks/grid_result.png"
	verticalModel    = "./mocks/gopher_aviator_vertical.png"
	horizontalModel  = "./mocks/gopher_aviator_horizontal.png"
	gridModel        = "./mocks/grid_model.png"
	fake_path        = "./fake-path.png"
	not_image_path   = "./README.md"
)

var concatifyTests = []concatifyTest{
	{
		"Concat two images vertical",
		[]string{gopher, aviator},
		verticalResult,
		verticalModel,
	},
	{
		"Concat two images horizontal",
		[]string{gopher, aviator},
		horizontalResult,
		horizontalModel,
	},
	{
		"Concat 6 files into 2x3 grid",
		[]string{aviator, aviator, gopher, aviator, aviator, gopher},
		gridResult,
		gridModel,
	},
}

var concatifyTestsFail = []concatifyTestFail{
	{
		concatifyTest{
			"Second file doesn't not exist",
			[]string{gopher, fake_path},
			verticalResult, verticalModel,
		},
		"Cannot open file " + fake_path,
	},
	{
		concatifyTest{
			"First file is not an image",
			[]string{not_image_path, aviator},
			horizontalResult, horizontalModel,
		},
		"Cannot decode file " + not_image_path,
	},
	{
		concatifyTest{
			"Not enough images to fill a grid",
			[]string{not_image_path, aviator},
			gridResult, gridModel,
		},
		"#sources = 2, 3x3 = 9",
	},
}

func TestNewVertical(t *testing.T) {
	test := concatifyTests[0]
	cimg, err := concatify.NewVertical(test.paths)
	if err != nil {
		t.Error(err)
	}
	cimg.Save(test.result)
	if !comparePNG(test.result, test.model) {
		t.Errorf("%s: result (%s) and model(%s) images are different!", test.name, test.result, test.model)
	}
	remove(test.result)
}

func TestNewHorizontal(t *testing.T) {
	test := concatifyTests[1]
	cimg, err := concatify.NewHorizontal(test.paths)
	if err != nil {
		t.Error(err)
	}
	cimg.Save(test.result)
	if !comparePNG(test.result, test.model) {
		t.Errorf("%s: result (%s) and model(%s) images are different!", test.name, test.result, test.model)
	}
	remove(test.result)
}

func TestNewGrid(t *testing.T) {
	test := concatifyTests[2]
	cimg, err := concatify.NewGrid(test.paths, 2, 3)
	if err != nil {
		t.Error(err)
	}
	cimg.Save(test.result)
	if !comparePNG(test.result, test.model) {
		t.Errorf("%s: result (%s) and model(%s) images are different!", test.name, test.result, test.model)
	}
	remove(test.result)
}

func TestDrawFail(t *testing.T) {
	for i, test := range concatifyTestsFail {
		var cimg *concatify.ConcatImage
		var err error
		switch i {
		case 0:
			cimg, err = concatify.NewVertical(test.paths)
		case 1:
			cimg, err = concatify.NewHorizontal(test.paths)
		case 2:
			cimg, err = concatify.NewGrid(test.paths, 3, 3)
		}
		_ = cimg
		if err != nil {
			if err.Error() != test.expectedErrorMessage {
				t.Errorf("%s: expected: %s, given: %s", test.name, test.expectedErrorMessage, err.Error())

			}
		}
	}
}

func BenchmarkDraw10(t *testing.B) {
	test := concatifyTests[1]
	for n := 0; n < 10; n++ {
		_, _ = concatify.NewVertical(test.paths)
		_, _ = concatify.NewHorizontal(test.paths)

	}
	remove(test.result)
}

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
