package array

import "testing"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] >= max {
			max = candies[i]
		}
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func Test_1431(t *testing.T) {
	t.Log(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) //[true true true false true]
	t.Log(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)) //[true false false false false]
	t.Log(kidsWithCandies([]int{12, 1, 12}, 10))    //[true false true]
}
