package main

import (
	"testing"
)

func TestTowSum(t *testing.T) {
	tests := [][]int{
		[]int{3, 2, 4}, //one
		[]int{1, 2, 3},
		[]int{1, 2},
		[]int{1, 5, 0, 15, 3, 2},
	}
	targets := []int{
		6,
		4,
		5,
		3,
	}
	results := [][]int{
		[]int{1, 2},
		[]int{0, 2},
		[]int{-1, -1},
		[]int{2, 4},
	}
	for i, _ := range tests {
		ret := TwoSum(tests[i], targets[i])
		if ret != nil && ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Fatalf("case %d failes :%v,\n", i, ret)
		}
	}

}
