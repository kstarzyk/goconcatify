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
	first, second, result, model string
	options                      ConcatedImageOptions
}

var TEST_GOLANG_PNG = "./mocks/golang.png"
var TEST_GOLANG_AVIATOR_PNG = "./mocks/golang-aviator.png"
var TEST_RESULT_VER_PNG = "./mocks/golang-and-golang-aviator-ver.png"
var TEST_RESULT_HOR_PNG = "./mocks/test-result-ver.png"
var TEST_MODEL_RESULT_VER_PNG = "./mocks/test-result-ver.png"
var TEST_MODEL_RESULT_HOR_PNG = "./mocks/test-result-hor.png"
var TEST_DEFAULT_PNG_OPTIONS = ConcatedImageOptions{VERTICAL, false, false}
var TEST_HOR_PNG_OPTIONS = ConcatedImageOptions{HORIZONTAL, false, false}

var concatifyTests = []concatifyTest{
	{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG, TEST_RESULT_VER_PNG, TEST_MODEL_RESULT_VER_PNG, TEST_DEFAULT_PNG_OPTIONS},
}

func TestDrawVertical(t *testing.T) {
	cimg := NewConcatedImage([]string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG})
	cimg.Draw(TEST_RESULT_VER_PNG)
	res := comparePNG(TEST_RESULT_VER_PNG, TEST_MODEL_RESULT_VER_PNG)
	if !res {
		t.Errorf("Test failed")
		remove(TEST_RESULT_VER_PNG)
	}

	remove(TEST_RESULT_VER_PNG)
}

func TestDrawHorizontal(t *testing.T) {
	cimg := NewConcatedImage([]string{TEST_GOLANG_PNG, TEST_GOLANG_AVIATOR_PNG}, TEST_HOR_PNG_OPTIONS)
	cimg.Draw(TEST_RESULT_HOR_PNG)
	res := comparePNG(TEST_RESULT_HOR_PNG, TEST_MODEL_RESULT_HOR_PNG)
	if !res {
		t.Errorf("Test failed")
	}

	remove(TEST_RESULT_HOR_PNG)
}
