package ri

func reverse(x int) int {
	inverse := false
	max := 2147483647

	if x < 0 {
		x = 0 - x
		inverse = true
	}
	num := 0
	for x%10 > 0 || x >= 10 {
		val := x % 10
		x = x / 10
		num = num*10 + val
		if num >= max && !inverse {
			return 0
		}
		if num >= max+1 && inverse {
			return 0
		}
	}
	if inverse {
		return 0 - num
	}
	return num
}
