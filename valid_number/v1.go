package vn

func isNumber(s string) bool {
	const (
		Default = "defualt"
		Digit   = "digit"
		Sign    = "sign"
		Blank   = "blank"
	)

	stateMap := []map[string]int{
		{Blank: 0, Digit: 3, Sign: 1, ".": 2}, //0. Blank
		{Digit: 3, ".": 2},                    //1. Sign
		{Digit: 5},                            //2. "."
		{Digit: 3, ".": 8, "e": 4, Blank: 9},  //3. Digit*
		{Digit: 6, Sign: 7},                   //4. "e"
		{Digit: 5, "e": 4, Blank: 9},          //5. Digit after "."*
		{Digit: 6, Blank: 9},                  //6. Digit after "e"*
		{Digit: 6},                            //7. Sign after "e"
		{Digit: 5, "e": 4, Blank: 9},          //8. "." after Digit*
		{Blank: 9},                            //9. End Blank*
	}

	state := 0
	for _, c := range []byte(s) {
		charType := Default
		if c >= '0' && c <= '9' {
			charType = Digit
		} else if c == ' ' {
			charType = Blank
		} else if c == '+' || c == '-' {
			charType = Sign
		} else {
			charType = string(c)
		}
		if ns, exists := stateMap[state][charType]; !exists {
			return false
		} else {
			state = ns
		}
	}
	if state == 8 || state == 3 || state == 5 || state == 6 || state == 9 {
		return true
	}
	return false
}
