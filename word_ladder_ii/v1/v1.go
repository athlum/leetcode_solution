package v1

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if len(beginWord) != len(endWord) {
		return [][]string{}
	}
	bl := len(beginWord)
	if bl == 0 {
		return [][]string{}
	}

	lm := map[string]bool{}
	findEnd := false
	for _, w := range wordList {
		if (!findEnd) && w == endWord {
			findEnd = true
		}
		if _, ok := lm[w]; ok {
			return [][]string{}
		} else {
			lm[w] = false
		}
	}
	if !findEnd {
		return [][]string{}
	}
	paths := map[string]bool{
		beginWord: false,
	}
	route := map[string][]string{}
	meetEnd := false
	for len(paths) > 0 && !meetEnd {
		for w, _ := range paths {
			delete(lm, w)
		}
		np := map[string]bool{}
		for w, _ := range paths {
			for i, _ := range w {
				for j := 0; j < 26; j += 1 {
					nw := w[:i] + string(rune('a'+j)) + w[i+1:]
					if _, ok := lm[nw]; ok {
						if _, ok := route[nw]; ok {
							route[nw] = append(route[nw], w)
						} else {
							route[nw] = []string{w}
						}
						np[nw] = false
					}
				}
			}
		}
		paths = np
		_, meetEnd = paths[endWord]
	}
	r := [][]string{}
	backTrace(endWord, beginWord, route, []string{endWord}, &r)
	return r
}

func backTrace(s, t string, route map[string][]string, cur []string, r *[][]string) {
	ps, ok := route[s]
	if !ok {
		return
	}
	for _, p := range ps {
		nc := append([]string{p}, cur...)
		if p == t {
			*r = append(*r, nc)
			return
		} else {
			backTrace(p, t, route, nc, r)
		}
	}
}
