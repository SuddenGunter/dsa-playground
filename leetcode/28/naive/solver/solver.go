package solver

func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			matches := true

			j := 1

			for j < len(needle) && i+j < len(haystack) {
				if haystack[i+j] != needle[j] {
					matches = false
				}

				j++
			}

			if matches && j == len(needle) {
				return i
			}
		}
	}

	return -1
}
