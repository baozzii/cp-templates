package templates

func primitive_root(m int) int {
	switch m {
	case 2:
		return 1
	case 167772161:
		return 3
	case 469762049:
		return 3
	case 754974721:
		return 11
	case 998244353:
		return 3
	default:
		var divs [20]int
		divs[0] = 2
		cnt := 1
		x := (m - 1) / 2
		for x%2 == 0 {
			x >>= 1
		}
		for i := 3; i*i <= x; i += 2 {
			if x%i == 0 {
				divs[cnt] = i
				cnt++
				for x%i == 0 {
					x /= i
				}
			}
		}
		if x > 1 {
			divs[cnt] = x
			cnt++
		}
		for g := 2; ; g++ {
			ok := true
			for i := range cnt {
				if pow(g, (m-1)/divs[i], m) == 1 {
					ok = false
					break
				}
			}
			if ok {
				return g
			}
		}
	}
}
