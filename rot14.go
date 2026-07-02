package piscine

func Rot14(s string) string {
	bs := []byte(s)
	var res []byte
	for _, v := range bs {
		if v >= 65 && v <= 90 {
			if v > 76 {
				v -= 12
			} else {
				v += 14
			}
			res = append(res, v)
		} else if v >= 97 && v <= 122 {
			if v > 108 {
				v -= 12
			} else {
				v += 14
			}
			res = append(res, v)
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
