package main

import "math"

//第一题
func maxDiff(arr []int64) int64 {

	max := math.MinInt64
	min := math.MaxInt64
	for v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return int64(max - min)
}
