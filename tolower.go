package piscine

func ToLower(s string) string {
	byted_s := []byte(s)
	for index, i := range byted_s {
		if i >= 65 && i <= 90 {
			byted_s[index] = byted_s[index] + 32
		} else {
			continue
		}
	}
	return string(byted_s)
}
