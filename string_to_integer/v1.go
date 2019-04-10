package sti

func myAtoi(str string) int {
	inverse := false
	start := false
	end := false
	num := 0
	max := 2147483647
	min := -2147483648
	for _, c := range []byte(str) {
		if c == ' ' {
			if start {
				break
			}
			continue
		} else if c == '+' && !start && !end {
			start = true
		} else if c == '-' && !start && !end {
			start = true
			inverse = true
		} else if c >= '0' && c <= '9' && !end {
			if !start {
				start = true
			}
			val := 0
			switch c {
			case '1':
				val = 1
			case '2':
				val = 2
			case '3':
				val = 3
			case '4':
				val = 4
			case '5':
				val = 5
			case '6':
				val = 6
			case '7':
				val = 7
			case '8':
				val = 8
			case '9':
				val = 9
			}
			num = num*10 + val
			if num >= max && !inverse {
				return max
			}
			if num >= max+1 && inverse {
				return min
			}
		} else if !start {
			return 0
		} else if start {
			break
		} else {
			return 0
		}
	}
	if inverse {
		num = 0 - num
	}
	return num
}
