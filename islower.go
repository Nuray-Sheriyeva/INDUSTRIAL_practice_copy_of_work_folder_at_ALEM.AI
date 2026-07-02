package piscine

func IsLower(s string) bool {
	byted_s := []byte(s)
	for _, i := range byted_s {
		if i >= 97 && i <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
