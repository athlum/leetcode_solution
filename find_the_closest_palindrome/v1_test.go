package ftcp

import (
	"fmt"
	"testing"
)

func TestCases(t *testing.T) {
	if v := nearestPalindromic("123"); v != "121" {
		t.Errorf("error: %v", v)
	}
	if v := nearestPalindromic("99"); v != "101" {
		t.Errorf("error: %v", v)
	}
	if v := nearestPalindromic("100"); v != "99" {
		t.Errorf("error: %v", v)
	}
	if v := nearestPalindromic("10"); v != "9" {
		t.Errorf("error: %v", v)
	}
	if v := nearestPalindromic("11"); v != "9" {
		t.Errorf("error: %v", v)
	}
}

func TestPrint(t *testing.T) {
	fmt.Println(cases("123", 3))
	fmt.Println(cases("99", 2))
	fmt.Println(cases("100", 3))
	fmt.Println(cases("10", 2))
	fmt.Println(cases("11", 2))
}
