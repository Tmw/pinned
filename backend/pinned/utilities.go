package pinned

func cap(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
