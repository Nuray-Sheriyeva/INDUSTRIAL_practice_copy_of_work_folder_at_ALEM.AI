package piscine

func IsNumeric(s string) bool {
	byted_s := []byte(s)
	for _, i := range byted_s {
		if i >= 48 && i <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
