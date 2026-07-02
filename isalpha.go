package piscine

func IsAlpha(s string) bool {
	byted_s := []byte(s)
	for _, i := range byted_s {
		if (i >= 97 && i <= 122) || (i >= 65 && i <= 90) || (i >= 48 && i <= 57) {
			continue
		} else {
			return false
		}
	}
	return true
}
