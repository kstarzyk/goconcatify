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

type concatifyTestFail struct {
	concatifyTest
	expectedErrorMessage string
}

var (
	TEST_GOLANG_PNG           = "./mocks/golang.png"
	TEST_GOLANG_AVIATOR_PNG   = "./mocks/golang-aviator.png"
	TEST_RESULT_VER_PNG       = "./mocks/golang-and-golang-aviator-ver.png"
	TEST_RESULT_HOR_PNG       = "./mocks/golang-ang-golang-aviator-hor.png"
	TEST_MODEL_RESULT_VER_PNG = "./mocks/test-result-ver.png"
	TEST_MODEL_RESULT_HOR_PNG = "./mocks/test-result-hor.png"
	TEST_FAKE_PATH            = "./fake-path.png"
	TEST_NOT_IMAGE            = "./README.md"
	TEST_DEFAULT_PNG_OPTIONS  = ConcatedImageOptions{VERTICAL, false, false}
	TEST_HOR_PNG_OPTIONS      = ConcatedImageOptions{HORIZONTAL, false, false}
)

var concatifyTests = []concatifyTest{
	{"Concat two files vertical (default settings)", []string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG}, TEST_RESULT_VER_PNG, TEST_MODEL_RESULT_VER_PNG, TEST_DEFAULT_PNG_OPTIONS},
	{"Concat two files horizontal", []string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG}, TEST_RESULT_HOR_PNG, TEST_MODEL_RESULT_HOR_PNG, TEST_HOR_PNG_OPTIONS},
}

var concatifyTestsFail = []concatifyTestFail{
	{concatifyTest{"Second file doesn't not exist", []string{TEST_GOLANG_PNG, TEST_FAKE_PATH}, TEST_RESULT_HOR_PNG, TEST_MODEL_RESULT_HOR_PNG, TEST_HOR_PNG_OPTIONS}, "Cannot open file " + TEST_FAKE_PATH},
	{concatifyTest{"First file is not an image", []string{TEST_NOT_IMAGE, TEST_GOLANG_AVIATOR_PNG}, TEST_RESULT_HOR_PNG, TEST_MODEL_RESULT_HOR_PNG, TEST_HOR_PNG_OPTIONS}, "Cannot decode file " + TEST_NOT_IMAGE},
}

func TestDraw(t *testing.T) {
	for _, test := range concatifyTests {
		cimg, err := NewConcatedImage(test.paths, test.options)
		if err != nil {
			t.Error(err)
		}
		cimg.Draw(test.result)
		if !comparePNG(test.result, test.model) {
			t.Errorf("%s: result (%s) and model(%s) images are different!", test.name, test.result, test.model)
		}
		remove(test.result)
	}
}

func TestDrawFail(t *testing.T) {
	for _, test := range concatifyTestsFail {
		cimg, err := NewConcatedImage(test.paths, test.options)
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
		cimg, _ := NewConcatedImage(test.paths)
		cimg.Draw(test.result)
	}
	remove(test.result)
}
