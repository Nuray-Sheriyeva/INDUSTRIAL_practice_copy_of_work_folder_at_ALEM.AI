package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
