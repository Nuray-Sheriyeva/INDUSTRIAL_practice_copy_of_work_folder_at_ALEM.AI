package piscine

func IsUpper(s string) bool {
	byted_s := []byte(s)
	for _, i := range byted_s {
		if i >= 65 && i <= 90 {
			continue
		} else {
			return false
		}
	}
	return true
}
