package whitestrings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSimilarity(t *testing.T) {
	var testcases = map[string]struct {
		inputStr1          string
		inputStr2          string
		expectedSimilarity float64
	}{
		"Good #1": {
			inputStr1:          "Healed",
			inputStr2:          "Sealed",
			expectedSimilarity: 0.8,
		},
		"Good #2": {
			inputStr1:          "Healed",
			inputStr2:          "Healthy",
			expectedSimilarity: 0.5454545454545454,
		},
		"Good #3": {
			inputStr1:          "Healed",
			inputStr2:          "Heard",
			expectedSimilarity: 0.4444444444444444,
		},
		"Good #4": {
			inputStr1:          "Healed",
			inputStr2:          "Herded",
			expectedSimilarity: 0.40,
		},
		"Good #5": {
			inputStr1:          "Healed",
			inputStr2:          "Help",
			expectedSimilarity: 0.25,
		},
		"Good #6": {
			inputStr1:          "Healed",
			inputStr2:          "Sold",
			expectedSimilarity: 0.0,
		},
		"Good #7": {
			inputStr1:          "GGGGG",
			inputStr2:          "GG",
			expectedSimilarity: 0.4,
		},
		"Good #8": {
			inputStr1:          "EVEN DEATH MAY DIE",
			inputStr2:          "DEATH",
			expectedSimilarity: 0.5333333333333333,
		},
		"Good #9": {
			inputStr1:          "EVEN DEATH MAY DIE",
			inputStr2:          "EVEN DEATH MAY LIE",
			expectedSimilarity: 0.9090909090909091,
		},
		"Good #10": {
			inputStr1:          "bunnyfoofoo",
			inputStr2:          "bunnyfoofoo",
			expectedSimilarity: 1.0,
		},
		"Good #11": {
			inputStr1:          "bunny foo foo",
			inputStr2:          "bunny foo foo",
			expectedSimilarity: 1.0,
		},
		"Good #12: One character against another": {
			inputStr1:          "A",
			inputStr2:          "A",
			expectedSimilarity: 1.0,
		},
		"Good #13: One character against two #1": {
			inputStr1:          "A",
			inputStr2:          "AB",
			expectedSimilarity: 0.5,
		},
		"Good #14: One character against two #2": {
			inputStr1:          "A",
			inputStr2:          "AA",
			expectedSimilarity: 1.0,
		},
		"Empty str1": {
			inputStr1:          "",
			inputStr2:          "Sold",
			expectedSimilarity: 0.0,
		},
		"Empty str2": {
			inputStr1:          "Sold",
			inputStr2:          "",
			expectedSimilarity: 0.0,
		},
		"Empty both": {
			inputStr1:          "",
			inputStr2:          "",
			expectedSimilarity: 0.0,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actualSimilarity := GetSimilarity(tc.inputStr1, tc.inputStr2)
			require.Equal(t, tc.expectedSimilarity, actualSimilarity)
		})
	}
}
