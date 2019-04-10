package v1

import (
	"testing"
)

func Test_V1(t *testing.T) {
	v := maxPoints([]Point{
		{1, 1},
	})
	if v != 1 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{0, 0},
		{0, 0},
	})
	if v != 2 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{0, 0},
		{1, 1},
		{0, 0},
		{1, 1},
		{1, 2},
	})
	if v != 4 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{1, 1},
		{2, 2},
		{3, 3},
	})
	if v != 3 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	})
	if v != 4 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{0, 0},
		{1, 1},
		{1, -1},
	})
	if v != 2 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{2, 3},
		{3, 3},
		{-5, 3},
	})
	if v != 3 {
		t.Error("failed")
	}

	v = maxPoints([]Point{
		{1, 1},
		{1, 1},
		{2, 3},
	})
	if v != 3 {
		t.Error("failed")
	}
}
