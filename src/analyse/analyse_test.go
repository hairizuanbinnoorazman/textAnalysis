package analyse

import (
	"reflect"
	"testing"
)

func TestAnalyzeTop10(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		text           string
		expectedOutput []Word
	}{
		{"Simple alphabet case", "a a a a a", []Word{{"a", 5}}},
		{"Check for alphabetical order", "alpha, beta, gamma google, apple, amazon, twitter, rainbow, " +
			"handphone, pixel",
			[]Word{
				{"alpha", 1},
				{"amazon", 1},
				{"apple", 1},
				{"beta", 1},
				{"gamma", 1},
				{"google", 1},
				{"handphone", 1},
				{"pixel", 1},
				{"rainbow", 1},
				{"twitter", 1},
			}},
		{"Check that input is put through lower case", "alpha, Beta, GAMMA google, apple, AMAZon, " +
			"twitter, rainbow, handphone, pixel",
			[]Word{
				{"alpha", 1},
				{"amazon", 1},
				{"apple", 1},
				{"beta", 1},
				{"gamma", 1},
				{"google", 1},
				{"handphone", 1},
				{"pixel", 1},
				{"rainbow", 1},
				{"twitter", 1},
			}},
		{"Check for 11 inputs", "a b c d e f g h i j a a a a a j j j i e e k",
			[]Word{
				{"a", 6},
				{"j", 4},
				{"e", 3},
				{"i", 2},
				{"b", 1},
				{"c", 1},
				{"d", 1},
				{"f", 1},
				{"g", 1},
				{"h", 1},
			}},
		{"Check for special symbols", "i'm max-mega unit-tests",
			[]Word{
				{"i'm", 1},
				{"max-mega", 1},
				{"unit-tests", 1},
			}},
	}

	for _, singleTestCase := range testCases {
		actualResult := TopNwords(singleTestCase.text, 10)
		if !reflect.DeepEqual(actualResult, singleTestCase.expectedOutput) {
			t.Error(singleTestCase.testCaseName, "failed. Expected:", singleTestCase.expectedOutput, "Actual:", actualResult)
		}
	}
}

func TestRemovePunctuation(t *testing.T) {
	testCases := []struct {
		testCaseName string
		text         string
		expectedText string
	}{
		{"Remove punctuation - clean", "a a a a a", "a a a a a"},
		{"Remove punctuation - scattered punctuation", "muscle. I have eaten already; the trials", "muscle I have eaten already the trials"},
		{"Remove punctuation - multiple cases of punctuation", "muscle.... See the heavens;;!!", "muscle   See the heavens  "},
	}
	for _, singleTestCase := range testCases {
		actualResult := removePuctuation(singleTestCase.text)
		if actualResult != singleTestCase.expectedText {
			t.Error(singleTestCase.testCaseName, "failed. Expected:", singleTestCase.expectedText, "Actual:", actualResult)
		}
	}
}
