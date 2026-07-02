package piscine

func ListSize(l *List) int {
	it := l.Head
	count := 0
	for it != nil {
		count++
		it = it.Next
	}
	return count
}
