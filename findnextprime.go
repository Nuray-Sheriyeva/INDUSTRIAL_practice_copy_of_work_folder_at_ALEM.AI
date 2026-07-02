package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	} else if IsPrime(nb) {
		return nb
	} else {
		if nb%2 == 0 {
			nb++
		}
		for i := nb; ; i += 2 {
			if IsPrime(i) {
				return i
			}
		}
	}
}
