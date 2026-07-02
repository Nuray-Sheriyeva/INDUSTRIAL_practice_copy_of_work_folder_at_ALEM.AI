package piscine

func Unmatch(a []int) int {
	var freq []int
	var vis []int
	cont := true
	for _, v := range a {
		cont = true
		for _, j := range vis {
			if v == j {
				cont = false
				break
			}
		}
		if !cont {
			continue
		}
		count := 0
		for _, n := range a {
			if v == n {
				count++
			}
		}
		freq = append(freq, count)
		vis = append(vis, v)
	}
	lost := -1
	for i, v := range freq {
		if v%2 != 0 {
			lost = vis[i]
			break
		}
	}
	return lost
}
