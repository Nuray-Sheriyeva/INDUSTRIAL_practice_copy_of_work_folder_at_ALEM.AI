package piscine

func StringToIntSlice(str string) []int {
	runed_s := []rune(str)
	var inted_s []int
	for _, v := range runed_s {
		inte := int(v)
		inted_s = append(inted_s, inte)
	}

	return inted_s
}

func bytesToExtendedString(bytes []byte) string {
	runes := make([]rune, len(bytes))
	for i, b := range bytes {
		runes[i] = rune(b) // Directly copies the 0-255 byte value to a 32-bit rune
	}
	return string(runes)
}
