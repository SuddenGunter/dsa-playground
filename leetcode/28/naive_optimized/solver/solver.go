package solver

//
//func StrStr(haystack, needle string) int {
//	if needle == "" {
//		return 0
//	}
//
//	if len(needle) > len(haystack) {
//		return -1
//	}
//
//	for i := 0; i < len(haystack); {
//		if haystack[i] == needle[0] {
//			matches := true
//
//			j := 1
//
//			for j < len(needle) && i+j < len(haystack) {
//				if haystack[i+j] != needle[j] {
//					matches = false
//				}
//
//				j++
//			}
//
//			if matches && j == len(needle) {
//				return i
//			} else {
//				// skip partial match (but need to consider i++ below)
//				if j > 1 {
//					i += max(j-1, 1)
//				}
//			}
//		} else {
//			i++
//		}
//
//	}
//
//	return -1
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		// early return when unchecked part of haystack isn't long enough to contain needle
		if len(needle) > len(haystack)-i {
			break
		}

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
