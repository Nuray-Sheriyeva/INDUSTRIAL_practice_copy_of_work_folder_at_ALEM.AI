package piscine

func ConcatParams(args []string) string {
	s := ""
	for i, v := range args {
		if i != len(args)-1 {
			s += v + "\n"
		} else {
			s += v
		}
	}
	return s
}
