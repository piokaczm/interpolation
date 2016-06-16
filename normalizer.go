package interpolation

func NormalizeMap(m map[int]int, max int, min int) map[int]int {
	for i, val := range m {
		if val > max {
			m[i] = max
		} else if val < min {
			m[i] = min
		}
	}
	return m
}
