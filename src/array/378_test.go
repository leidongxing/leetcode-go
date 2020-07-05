package array

import (
	"sort"
	"testing"
)

//有序矩阵中第K小的元素
func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	nums := make([]int, m*n)
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			nums[index] = matrix[i][j]
			index++
		}
	}
	sort.Ints(nums)
	return nums[k-1]
}
func Test_378(t *testing.T) {
	t.Log(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
}
