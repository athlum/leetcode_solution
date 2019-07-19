package ftcp

// https://leetcode.com/problems/find-the-closest-palindrome/

import (
	"math"
	"strconv"
)

func reverse(n string) string {
	l := len(n)
	v := make([]byte, l)
	for i, c := range []byte(n) {
		v[l-1-i] = c
	}
	return string(v)
}

func mirror(left string, shift int) string {
	leftm := reverse(left)

	return left[:len(left)-shift] + leftm
}

func cases(n string, total int) []string {
	half := total / 2
	odd := total%2 > 0
	shift := 0
	if odd {
		half += 1
		shift += 1
	}
	left := n[:half]
	leftV, _ := strconv.Atoi(left)

	oriMirror := mirror(left, shift)

	leftPlus := strconv.Itoa(leftV + 1)
	pShift := shift
	if len(leftPlus) > len(left) {
		pShift += 1
	}
	plusMirror := mirror(leftPlus, pShift)

	leftMinus := strconv.Itoa(leftV - 1)
	mShift := shift
	if len(leftMinus) < len(left) {
		leftMinus += "9"
		mShift += 1
	} else if leftMinus == "0" {
		leftMinus = "9"
		mShift += 1
	}
	minusMirror := mirror(leftMinus, mShift)
	return []string{oriMirror, plusMirror, minusMirror}
}

func nearestPalindromic(n string) string {
	total := len(n)
	nv, _ := strconv.Atoi(n)
	if nv == 0 {
		return "1"
	} else if nv < 10 {
		return strconv.Itoa(nv - 1)
	}
	min := 0.0
	output := 0
	for _, s := range cases(n, total) {
		v, _ := strconv.Atoi(s)
		if dis := math.Abs(float64(v - nv)); min == 0.0 || dis < min {
			min = dis
			output = v
		} else if dis == min && v < output {
			output = v
		}
	}
	return strconv.Itoa(output)
}
