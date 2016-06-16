package interpolate

func normalizeMap(m map[int]int, max int, min int) map[int]int {
	for i, pair := range m {
		if rate > max {
			rates[i] = max
		} else if rate < min {
			rates[i] = min
		}
	}
	return rates
}
