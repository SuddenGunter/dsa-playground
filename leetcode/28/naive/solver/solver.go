package solver

func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		// early return when unchecked part of haystack isn't long enough to contain needle
		if len(needle) > len(haystack)-i {
			break
		}

		if haystack[i] == needle[0] {
			matches := true

			j := 1

			for j < len(needle) {
				if haystack[i+j] != needle[j] {
					matches = false
					break
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
