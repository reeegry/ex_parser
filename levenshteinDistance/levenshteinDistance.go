package levenshteinDistance

import "github.com/agnivade/levenshtein"

func FindDistance(text1ptr *string, text2ptr *string) int {
	text1 := *text1ptr
	text2 := *text2ptr
	return levenshtein.ComputeDistance(text1, text2)
}
