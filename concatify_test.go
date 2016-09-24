package concatify

import (
	"os"
	"testing"
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
	options       ConcatedImageOptions
}

var (
	TEST_GOLANG_PNG           = "./mocks/golang.png"
	TEST_GOLANG_AVIATOR_PNG   = "./mocks/golang-aviator.png"
	TEST_RESULT_VER_PNG       = "./mocks/golang-and-golang-aviator-ver.png"
	TEST_RESULT_HOR_PNG       = "./mocks/golang-ang-golang-aviator-hor.png"
	TEST_MODEL_RESULT_VER_PNG = "./mocks/test-result-ver.png"
	TEST_MODEL_RESULT_HOR_PNG = "./mocks/test-result-hor.png"
	TEST_DEFAULT_PNG_OPTIONS  = ConcatedImageOptions{VERTICAL, false, false}
	TEST_HOR_PNG_OPTIONS      = ConcatedImageOptions{HORIZONTAL, false, false}
)

var concatifyTests = []concatifyTest{
	{"Concat two files vertical (default settings)", []string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG}, TEST_RESULT_VER_PNG, TEST_MODEL_RESULT_VER_PNG, TEST_DEFAULT_PNG_OPTIONS},
	{"Concat two files horizontal", []string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG}, TEST_RESULT_HOR_PNG, TEST_MODEL_RESULT_HOR_PNG, TEST_HOR_PNG_OPTIONS},
}

func TestDraw(t *testing.T) {
	for _, test := range concatifyTests {
		cimg := NewConcatedImage(test.paths, test.options)
		cimg.Draw(test.result)
		if !comparePNG(test.result, test.model) {
			t.Errorf("%s: result (%s) and model(%s) images are different!", test.name, test.result, test.model)
		}
		remove(test.result)
	}
}
