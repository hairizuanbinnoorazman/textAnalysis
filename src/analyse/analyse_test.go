package analyse

import (
	"reflect"
	"testing"
)

func TestAnalyzeTop10(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		text           string
		expectedOutput map[string]int
	}{
		{"Simple alphabet case", "a a a a a", map[string]int{"a": 5}},
		{"Check for alphabetical order", "alpha, beta, gamma google, apple, amazon, twitter, rainbow, " +
			"handphone, pixel",
			map[string]int{
				"alpha":     1,
				"amazon":    1,
				"apple":     1,
				"beta":      1,
				"gamma":     1,
				"google":    1,
				"handphone": 1,
				"pixel":     1,
				"rainbow":   1,
				"twitter":   1,
			}},
		{"Check that input is put through lower case", "alpha, Beta, GAMMA google, apple, AMAZon, " +
			"twitter, rainbow, handphone, pixel",
			map[string]int{
				"alpha":     1,
				"amazon":    1,
				"apple":     1,
				"beta":      1,
				"gamma":     1,
				"google":    1,
				"handphone": 1,
				"pixel":     1,
				"rainbow":   1,
				"twitter":   1,
			}},
		{"Check for cases of multiple inputs", "a b c d e f g h i j a a a a a j j j i e e",
			map[string]int{
				"a": 6,
				"j": 4,
				"e": 3,
				"i": 2,
				"b": 1,
				"c": 1,
				"d": 1,
				"f": 1,
				"g": 1,
				"h": 1,
			}},
	}

	for _, singleTestCase := range testCases {
		actualResult := top10words(singleTestCase.text)
		if !reflect.DeepEqual(actualResult, singleTestCase.expectedOutput) {
			t.Error(singleTestCase.testCaseName, "failed. Expected:", singleTestCase.expectedOutput, "Actual:", actualResult)
		}
	}
}
