package piscine

func AtoiBase(s string, base string) int {
	// Base must contain at least 2 characters
	if len(base) < 2 {
		return 0
	}
	// Validate base rules and map characters to their values
	baseMap := make(map[rune]int)
	for i, ch := range base {
		// Base cannot contain '+' or '-' characters
		if ch == '+' || ch == '-' {
			return 0
		}
		// Each character in the base must be unique
		if _, exists := baseMap[ch]; exists {
			return 0
		}
		baseMap[ch] = i
	}
	// Convert string 's' using the verified base
	result := 0
	baseLength := len(base)
	for _, ch := range s {
		val, exists := baseMap[ch]
		if !exists {
			return 0
		}
		result = result*baseLength + val
	}

	return result
}
