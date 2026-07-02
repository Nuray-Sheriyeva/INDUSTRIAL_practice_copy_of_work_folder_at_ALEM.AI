package piscine

func Capitalize(s string) string {
	byted_s := []byte(s)
	if len(s) == 1 {
		if byted_s[0] >= 97 && byted_s[0] <= 122 {
			byted_s[0] = byted_s[0] - 32
		}
	}
	for index, letter := range byted_s {
		if index != len(s)-1 {
			if index == 0 && letter >= 97 && letter <= 122 {
				byted_s[index] = byted_s[index] - 32
				if byted_s[index+1] >= 65 && byted_s[index+1] <= 90 {
					byted_s[index+1] = byted_s[index+1] + 32
				} else {
					continue
				}
			} else if !((letter >= 97 && letter <= 122) || (letter >= 65 && letter <= 90) || (letter >= 48 && letter <= 57)) {
				if byted_s[index+1] >= 97 && byted_s[index+1] <= 122 {
					byted_s[index+1] = byted_s[index+1] - 32
				} else {
					continue
				}
			} else if letter >= 97 && letter <= 122 {
				if byted_s[index+1] >= 65 && byted_s[index+1] <= 90 {
					byted_s[index+1] = byted_s[index+1] + 32
				} else {
					continue
				}
			} else if (letter >= 65 && letter <= 90) || (letter >= 48 && letter <= 57) {
				if byted_s[index+1] >= 65 && byted_s[index+1] <= 90 {
					byted_s[index+1] = byted_s[index+1] + 32
				} else {
					continue
				}
			}
		}
	}
	return string(byted_s)
}
