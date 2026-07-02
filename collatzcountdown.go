package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	} else if start == 1 {
		return 0
	} else {
		steps := 0
		if start%2 == 0 {
			start = start / 2
			steps += 1
		} else {
			start = 3*start + 1
			steps += 1
		}
		return steps + CollatzCountdown(start)
	}
}
