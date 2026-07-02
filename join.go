package piscine

func Join(elems []string, sep string) string {
	s := ""
	for i := 0; i < len(elems); i++ {
		if i == len(elems)-1 {
			s += elems[i]
		} else {
			s += elems[i] + sep
		}
	}
	return s
}
