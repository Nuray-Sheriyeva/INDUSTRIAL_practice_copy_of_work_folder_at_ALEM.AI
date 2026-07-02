package piscine

func AlphaCount(s string) int {
	counter := 0
	byted_string := []byte(s)
	for _, letter := range byted_string {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			counter++
		}
	}
	return counter
}
