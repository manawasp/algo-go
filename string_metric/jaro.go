package sdist

import (
	"math"
)

func Jaro(s1, s2 string) float64 {
	var matches, transposition float64
	len1 := float64(len(s1))
	len2 := float64(len(s2))
	matchRange := int(math.Floor(math.Max(len1, len2)/2)) - 1
	min := int(math.Min(len1, len2))

	if matchRange < 0 {
		matchRange = 0
	}

	for i := 0; i < min; i++ {
		if s1[i] == s2[i] {
			matches++
			continue
		}
		for j := 1; j <= matchRange; j++ {
			p := i + j
			if p < len(s1) && s1[p] == s2[i] &&
				p < len(s2) && s2[p] == s1[i] {
				matches += 2
				transposition++
			} else if p < len(s1) && s1[p] == s2[i] ||
				p < len(s2) && s2[p] == s1[i] {
				matches++
			}
		}
	}

	if matches == 0 {
		return 0
	}

	return ((matches / len1) + (matches / len2) + ((matches - transposition) / matches)) / 3
}
