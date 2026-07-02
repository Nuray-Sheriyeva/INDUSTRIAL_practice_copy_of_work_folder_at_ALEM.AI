package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result []bool
	for _, v := range tab {
		result = append(result, f(v))
	}
	count := 0
	for _, v := range result {
		if v {
			count++
		}
	}
	return count
}
