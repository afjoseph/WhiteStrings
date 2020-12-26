package whitestrings

import (
	"strings"
)

// getIndexOfElemInStringSlice returns the index of 'needle' in the string
// slice 'haystack'.  It returns -1 if no value was found.
func getIndexOfElemInStringSlice(haystack []string, needle string) int {
	for k, v := range haystack {
		if needle == v {
			return k
		}
	}
	return -1
}

// findWordPairs takes a string 'str' and returns a slice of 2-character
// strings (a pair) of each word, while ignoring all whitespaces and
// capitlizing 'str'.
//
// For example: it will turn 'bunnyfoofoo' to
// ["BU","UN","NN","NY","YF","FO","OO","OF","FO","OO"]
func findWordPairs(str string) []string {
	var wordPairs []string
	wordArr := strings.Split(str, " ")
	for _, word := range wordArr {
		word = strings.ToUpper(word)
		for pos := 0; pos < len(word); pos++ {
			if (pos + 1) != len(word) {
				wordPairs = append(wordPairs, string(word[pos:pos+2]))
			}
		}
	}
	return wordPairs
}

// getSimilarityWithOneCharString loops over 'other' and counts how many times
// did it see 'shortest' in it, then divides by the number of occurrences
// against the length of other
func getSimilarityWithOneCharString(shortest, other string) float64 {
	if len(shortest) != 1 {
		panic("Unexpected situation")
	}

	numOfOccurrences := 0
	for i := 0; i < len(other); i++ {
		if shortest[0] == other[i] {
			numOfOccurrences++
		}
	}
	return float64(numOfOccurrences) / float64(len(other))
}

// GetSimilarity runs the 'Strike A Match' algorithm from Simon White of
// Catalysoft (http://www.catalysoft.com/articles/StrikeAMatch.html) against
// 'str1' and 'str2'.
//
// NOTE: The original algorithm doesn't handle 1-character strings. I'm
// handling it in getSimilarityWithOneCharString() using a simple counter. See
// the function for more info.
//
// NOTE: This algorithm is case-insensitive.
func GetSimilarity(str1, str2 string) float64 {
	if len(str1) == 0 || len(str2) == 0 {
		return 0.0
	}
	if len(str1) == 1 {
		return getSimilarityWithOneCharString(str1, str2)
	}
	if len(str2) == 1 {
		return getSimilarityWithOneCharString(str2, str1)
	}

	pairs1 := findWordPairs(str1)
	pairs2 := findWordPairs(str2)
	union := len(pairs1) + len(pairs2)

	intersection := 0.0
	for _, pair1 := range pairs1 {
		index := getIndexOfElemInStringSlice(pairs2, pair1)
		if index == -1 {
			continue
		}
		intersection += 1
		// Remove element at index
		pairs2 = append(pairs2[:index], pairs2[index+1:]...)
	}

	return (2.0 * intersection) / float64(union)
}
