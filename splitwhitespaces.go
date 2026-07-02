package piscine

func SplitWhiteSpaces(s string) []string {
	byted_s := []byte(s)
	new := ""
	var array []string
	for _, v := range byted_s {
		if v == 32 || v == 9 || v == 10 {
			if new != "" {
				array = append(array, new)
				new = ""
			}
		} else {
			new += string(v)
		}
	}
	if new != "" {
		array = append(array, new)
	}
	return array
}
