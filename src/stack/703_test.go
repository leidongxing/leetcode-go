package stack

import "testing"

//数据流中的第K大元素
type KthLargest struct {
}

func Constructor_703(k int, nums []int) KthLargest {

}

func (this *KthLargest) Add(val int) int {

}

func Test_703(t *testing.T) {
	k := 3
	arr := []int{4, 5, 8, 2}
	Constructor_703(k, arr)
	KthLargest.Add(3)  // returns 4
	KthLargest.Add(5)  // returns 5
	KthLargest.Add(10) // returns 5
	KthLargest.Add(9)  // returns 8
	KthLargest.Add(4)  // returns 8
}
