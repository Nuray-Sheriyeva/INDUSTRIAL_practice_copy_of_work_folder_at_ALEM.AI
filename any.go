package piscine

func Any(f func(string) bool, a []string) bool {
	var result []bool
	for _, v := range a {
		result = append(result, f(v))
	}
	for _, v := range result {
		if v {
			return true
		}
	}
	return false
}
