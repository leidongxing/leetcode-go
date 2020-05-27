package greedy

import (
	"sort"
	"testing"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	max := 0
	answer := 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		answer += satisfaction[i]
		if answer < 0 {
			break
		}
		//每加一次 都会把之前的数再多加一遍
		max += answer
	}
	return max
}

func Test_1402(t *testing.T) {
	//t.Log(maxSatisfaction([]int{-1, -8, 0, 5, -9}))    //14   -9 -8 -1 0 5
	//t.Log(maxSatisfaction([]int{4, 3, 2}))             //20
	//t.Log(maxSatisfaction([]int{-1, -4, -5}))          //0
	t.Log(maxSatisfaction([]int{-2, 5, -1, 0, 3, -3})) //35  -3 -2 -1 0 3 5   -3-4-3+15+30
}
