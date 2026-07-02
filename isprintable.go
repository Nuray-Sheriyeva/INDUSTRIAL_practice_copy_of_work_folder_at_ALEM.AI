package piscine

func IsPrintable(s string) bool {
	byted_s := []byte(s)
	for _, i := range byted_s {
		if i <= 31 || i == 127 {
			return false
		} else {
			continue
		}
	}
	return true
}
