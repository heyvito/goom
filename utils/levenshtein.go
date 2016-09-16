package utils

import "unicode/utf8"

// Levenshtein calculates the difference between two strings by using the
// Levenshtein algorithm
func Levenshtein(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)

	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0] // fj1 has the value of f[j - 1], of last interaction
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1) // delete and insert
			if cb != ca {
				mn = min(mn, fj1+1) // change
			} else {
				mn = min(mn, fj1) // matched
			}

			fj1, f[j] = f[j], mn // save f[j] to fj1 as j is about to increase, update f[j] to mn
			j++
		}
	}

	return f[len(f)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
