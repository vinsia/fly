package fly

func FuzzySearch(filterValue string, optValue string, optIndex int) bool {
	// check filterValue is optValue subsequence
	// e.g. "guangzhou002" is subsequence of "guangzhou-server-002"
	next := 0
	for i := 0; i < len(filterValue); i += 1 {
		match := false
		for next < len(optValue) {
			if optValue[next] == filterValue[i] {
				match = true
				next += 1
				break
			}
			next += 1
		}
		if !match {
			return false
		}
	}
	return true
}
