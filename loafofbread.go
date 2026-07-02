package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		ns := ""
		count := 0
		for i := 0; i < len(str); i++ {
			if count != 5 {
				if str[i] != ' ' {
					ns += string(str[i])
					count++
				}
			} else {
				if i != len(str)-1 {
					ns += " "
					count = 0
					continue
				}
			}
		}
		return ns + "\n"
	}
}
