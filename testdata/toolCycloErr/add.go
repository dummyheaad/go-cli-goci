package add

func add(a, b int) int {
	result := 0

	if a > 0 {
		if b > 0 {
			for i := 0; i < a; i++ {
				result++
			}
			for j := 0; j < b; j++ {
				result++
			}
		} else if b < 0 {
			for i := 0; i < a; i++ {
				result++
			}
			for j := 0; j > b; j-- {
				result--
			}
		} else {
			return a
		}
	} else if a < 0 {
		if b > 0 {
			for i := 0; i > a; i-- {
				result--
			}
			for j := 0; j < b; j++ {
				result++
			}
		} else if b < 0 {
			for i := 0; i > a; i-- {
				result--
			}
			for j := 0; j > b; j-- {
				result--
			}
		} else {
			return a
		}
	} else {
		return b
	}

	return result
}
