package piscine

func SortWordArr(a []string) []string {
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
