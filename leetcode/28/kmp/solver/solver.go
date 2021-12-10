package solver

func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	prefixArray := make([]int, len(needle))
	prefixArray[0] = 0

	for i, j := 1, 0; i < len(needle); {
		if needle[i] == needle[j] {
			j++
			prefixArray[i] = j
			i++
		} else {
			if j != 0 {
				j = prefixArray[j-1]
			} else {
				prefixArray[i] = 0
				i++
			}
		}
	}

	for i, j := 0, 0; i < len(haystack); {
		if needle[j] == haystack[i] {
			i++
			j++
		}

		if j == len(needle) {
			return i - j
		}

		if i < len(haystack) && needle[j] != haystack[i] {
			if j != 0 {
				j = prefixArray[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}
