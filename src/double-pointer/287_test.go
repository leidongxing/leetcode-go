package double_pointer

import "testing"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	//如果数组中存在相同的数 必然会跳出循环
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	//跳出循环  相遇点到环入口的距离等于起点到环入口的距离
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func Test_287(t *testing.T) {
	var nums = []int{1, 3, 4, 2, 2}
	t.Log(findDuplicate(nums))
}
