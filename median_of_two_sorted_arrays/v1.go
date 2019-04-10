package motsa

import "sort"

type nums []int

func (n nums) Len() int           { return len(n) }
func (n nums) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nums) Less(i, j int) bool { return n[i] > n[j] }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tl := len(nums1) + len(nums2)
	mid := tl / 2
	avg := true
	if tl%2 > 0 {
		mid = (tl - 1) / 2
		avg = false
	}
	base := nums1
	tar := nums2
	if len(nums1) < mid {
		base = nums2
		tar = nums1
	}

	re := make([]int, mid+1)
	pb := 0
	pt := 0
	for i := 0; i <= mid; i += 1 {
		if pb < len(base) && (pt >= len(tar) || base[pb] <= tar[pt]) {
			re[i] = base[pb]
			pb += 1
		} else {
			re[i] = tar[pt]
			pt += 1
		}
	}

	sort.Sort(nums(re))
	if avg {
		return float64(re[0]+re[1]) / 2.0
	}
	return float64(re[0])
}
